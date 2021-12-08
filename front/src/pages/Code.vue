<template>
  <q-page class="flex">
    <div class="row fit items-start justify-center">

      <div class="col-12 col-md-9">
        <q-card>
          <q-card-section style="overflow-wrap: break-word;">
            <div class="text-h6" v-if="onchain">
              Contract {{ extAddr }}
              <a target="_blank" v-if="explorerLink" :href="explorerLink">
                <q-icon name="open_in_new" />
                <q-tooltip :offset="[10,10]">View it on {{ explorerName }}</q-tooltip>
              </a>
            </div>
            <div class="text-h6" v-else>Binary code</div>

            <div class="text-subtitle2">
              <span v-if="extCodeID">code: {{ extCodeID }}</span>
              <span v-else-if="errorCodeID" class="text-negative">
                Error: {{ errorCodeID }}
              </span>
              <span v-else><q-skeleton type="text" /></span>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Code panel -->
      <div class="col-12 col-md-9" v-if="errorCodeView">
        <div class="row justify-center items-center">
          <div class="col-12 q-pt-md">
            <p class="text-h6 text-center text-negative">
              <q-icon name="error"/> Error: {{ errorCodeView }}
            </p>
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
                &nbsp;
                <a target="_blank" :href="bloxyFunctionLink(func.hash)">
                  <pre><q-icon name="open_in_new" /></pre>
                  <q-tooltip :offset="[0,0]">View it on Bloxy</q-tooltip>
                </a>
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
pre {
  margin: 0.3em;
}
</style>

<script>
import AnsiConverter from 'ansi-to-html';
import axios from 'axios';
import { getAxiosError } from '../util.js';
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

      extAddr: "",

      extCodeID: "",
      errorCodeID: null,

      codeLoaded: false,
      code: {
        asm: "",
        pseudocode: "",
        functions: [],
      },
      errorCodeView: null,
    };
  },
  created() {
    var vm = this;
    if (vm.onchain) {
      vm.extAddr = this.$route.params.id;
    }
    else {
      vm.extCodeID = this.$route.params.id;
    }

    var codePromise;
    if (vm.onchain) {
      codePromise = axios.get('/api/addr/' + vm.extAddr)
        .then(function(res) {
          vm.extCodeID = res.data.extCodeID;
          return vm.extCodeID;
        })
        .catch(function(err) {
          vm.errorCodeID = getAxiosError(err);
          console.log(vm.errorCodeID);
          vm.errorCodeView = "No code to decompile";
          console.error(err);
          // Explicitly stop this promise chain
          return Promise.reject(err);
        });
    } else {
      codePromise = Promise.resolve(vm.extCodeID);
    }

    codePromise
      .then((ecid) => axios.get('/api/deco/' + ecid))
      .then(function(res) {
        vm.code = res.data.contract;
        vm.codeLoaded = true;
      })
      .catch(function(err) {
        vm.errorCodeView = getAxiosError(err);
        console.error(err);
      });
  },
  computed: {
    explorerLink() {
      return Networks.addrExplorer(this.extAddr);
    },
    explorerName() {
      return Networks.explorerName(this.extAddr);
    },
    pseudocodeHtml() {
      return convertAnsi(this.code.pseudocode)
    },
  },
  methods: {
    functionNameHtml(func) {
      return convertAnsi(func.color_name)
    },
    bloxyFunctionLink(funcHash) {
      if (funcHash.startsWith("0x")) {
        funcHash = funcHash.substring(2);
      }
      return "https://bloxy.info/functions/" + funcHash;
    },
  },
}
</script>

