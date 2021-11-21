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
      <div class="col-12">
        <q-tabs v-model="tabCode" dense inline-label align="justify">
          <q-tab name="asm" icon="source" label="Assembly"></q-tab>
          <q-tab name="pseudo" icon="code" label="Pseudocode"></q-tab>
          <q-tab name="function" icon="list" label="Functions"></q-tab>
        </q-tabs>

        <q-separator></q-separator>

        <q-tab-panels v-model="tabCode" class="q-pb-xl">
          <q-tab-panel name="asm">
            <pre v-html="contract.asm"></pre>
          </q-tab-panel>
          <q-tab-panel name="pseudo">
            <pre v-html="pseudocodeHtml()"></pre>
          </q-tab-panel>
          <q-tab-panel name="function">
            <q-list dense separator>
              <q-item v-for="func in contract.functions" :key="func.hash">
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
      contract: {
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

    axios.get('/api/deco/' + vm.extendedCodeHash)
      .then(function(res) {
        var resJson = res.data.default;
        vm.contract = resJson.contract;
      })
      .catch(function(err) {
        vm.error = err;
        console.log(err);
      });
  },
  methods: {
    pseudocodeHtml() {
      var vm = this;
      var ansi = vm.contract.pseudocode;
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

