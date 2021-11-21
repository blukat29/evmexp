<template>
  <q-page class="flex">
    <div class="row fit items-start justify-center">

      <div class="col-12 col-md-9">
        <q-card>
          <q-card-section style="overflow-wrap: break-word;">
            <div class="text-h6" v-if="onchain">
              Contract {{ extendedAddr }}
              <a target="_blank" v-if="explorerLink" :href="explorerLink">
                <q-icon name="open_in_new" />
                <q-tooltip :offset="[10,10]">View it on {{ explorerName }}</q-tooltip>
              </a>
            </div>
            <div class="text-h6" v-else>Binary code</div>

            <div class="text-subtitle2">
              <span v-if="extendedCodeHash">code: {{ extendedCodeHash }}</span>
              <span v-else-if="errorCodeHash" class="text-negative">Error: {{ errorCodeHash }}</span>
              <span v-else><q-skeleton type="text" /></span>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Code panel -->
      <div class="col-12 col-md-9" v-if="errorCodeHash">
        <div class="row justify-center items-center">
          <div class="col-12 q-pt-md">
            <p class="text-h5 text-center text-negative"><q-icon name="error"/> Failed</p>
          </div>
        </div>
      </div>
      <div class="col-12 col-md-9" v-else-if="!codeLoaded">
        <q-skeleton height="400px" square />
      </div>
      <div class="col-12 col-md-9 q-pt-md" v-else>
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
            <pre v-html="pseudocodeHtml"></pre>
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

    <q-page-sticky position="bottom" :offset="[0,18]" v-if="codeLoaded">
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
import Networks from '../networks';

function convertAnsi(ansi) {
  var ansiConverter = new AnsiConverter({
    newline: true,
  });
  return ansiConverter.toHtml(ansi);
}

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
      errorCodeHash: "",
      codeLoaded: false,
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
          vm.extendedCodeHash = res.data.extendedCodeHash;
          return vm.extendedCodeHash;
        })
        .catch(function(err) {
          var res = err.response;
          if (res.status == 404) {
            vm.errorCodeHash = res.data.error;
          }
        });
    } else {
      codePromise = Promise.resolve(vm.extendedCodeHash);
    }

    codePromise
      .then((ech) => axios.get('/api/deco/' + ech))
      .then(function(res) {
        vm.code = res.data.contract;
        vm.codeLoaded = true;
      })
      .catch(function(err) {
        vm.error = err;
        console.log(err);
      });
  },
  computed: {
    explorerLink() {
      return Networks.addrExplorer(this.extendedAddr);
    },
    explorerName() {
      return Networks.explorerName(this.extendedAddr);
    },
    pseudocodeHtml() {
      return convertAnsi(this.code.pseudocode)
    },
  },
  methods: {
    functionNameHtml(func) {
      return convertAnsi(func.color_name)
    },
  },
}
</script>

