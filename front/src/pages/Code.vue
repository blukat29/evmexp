<template>
  <q-page class="flex flex-center">
    <div class="row fit items-start justify-center" style="max-width: 1200px">

      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="text-h6" v-if="onchain">Contract {{ extendedAddr }}</div>
            <div class="text-h6" v-else>Binary code</div>
            <div class="text-subtitle2">
              <span v-if="extendedCodeHash">{{ extendedCodeHash }}</span>
              <span v-else><q-skeleton type="text" /></span>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Code panel -->
      <div class="col-12" v-if="!code.asm">
        <q-skeleton height="400px" square />
      </div>
      <div class="col-12" v-else>
        <q-tabs v-model="tabCode" dense inline-label no-caps align="justify">
          <q-tab name="asm" icon="source" label="Assembly"></q-tab>
          <q-tab name="pseudo" icon="code" label="Pseudocode"></q-tab>
          <q-tab name="function" icon="list" label="Functions"></q-tab>
        </q-tabs>

        <q-separator></q-separator>

        <q-tab-panels v-model="tabCode" class="q-pb-xl">
          <q-tab-panel name="asm">
            <pre v-html="code.asm"></pre>
          </q-tab-panel>
          <q-tab-panel name="pseudo">
            <pre v-html="pseudocodeHtml()"></pre>
          </q-tab-panel>
          <q-tab-panel name="function">
            <q-list dense separator>
              <q-item v-for="func in code.functions" :key="func.hash">
                <pre v-html="functionNameHtml(func)"></pre>
                <pre v-if="func.payable" class="text-accent">&nbsp;payable</pre>
                <pre v-if="!!func.getter" class="text-accent">&nbsp;view</pre>
                <pre v-html="' @ ' + func.hash" class="text-info"></pre>
              </q-item>
            </q-list>
          </q-tab-panel>
        </q-tab-panels>
      </div>
    </div>

    <q-page-sticky position="bottom" :offset="[0,18]">
      <q-list bordered class="bg-grey-5">
        <!-- Query storage -->
        <q-expansion-item group="inspect" padding icon="explore" label="Inspect on-chain data">
          <q-card>
            <q-card-section>
              Not ready yet
            </q-card-section>
          </q-card>
        </q-expansion-item>
      </q-list>
    </q-page-sticky>
  </q-page>
</template>

<style>
</style>

<script>
import AnsiConverter from 'ansi-to-html';
import axios from 'axios';

export default {
  name: 'Code',
  props: {
    onchain: Boolean,
  },
  data() {
    return {
      tabCode: "function",
      extendedAddr: "",
      extendedCodeHash: "",
      code: {
        asm: "",
        pseudocode: "",
        functions: [],
      },
    };
  },
  created() {
    var vm = this;
    if (vm.onchain) {
      vm.extendedAddr = this.$route.params.id;
    }
    else {
      vm.extendedCodeHash = this.$route.params.id;
    }

    var codePromise;
    if (vm.onchain) {
      codePromise = axios.get('/api/addr/' + vm.extendedAddr)
        .then(function(res) {
          console.log(res);
          var resJson = res.data;
          console.log(resJson);
          vm.extendedCodeHash = resJson.extendedCodeHash;
          return vm.extendedCodeHash;
        });
    } else {
      codePromise = Promise.resolve(vm.extendedCodeHash);
    }

    codePromise
      .then((ech) => axios.get('/api/deco/' + ech))
      .then(function(res) {
        var resJson = res.data;
        vm.code = resJson.contract;
      })
      .catch(function(err) {
        vm.error = err;
        console.log(err);
      });
  },
  methods: {
    pseudocodeHtml() {
      var vm = this;
      var ansi = vm.code.pseudocode;
      var ansiConverter = new AnsiConverter({
        newline: true,
      });
      var html = ansiConverter.toHtml(ansi);
      return html;
    },
    functionNameHtml(func) {
      var ansi = func.color_name;
      var ansiConverter = new AnsiConverter({
        newline: true,
      });
      var html = ansiConverter.toHtml(ansi);
      return html;
    },
  },
}
</script>

