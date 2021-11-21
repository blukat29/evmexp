<template>
  <q-page class="flex flex-center">
    <div class="row">
      <div class="col-12">
        <q-card>
          <q-tabs no-caps v-model="mode" align="justify"
                  class="text-grey" active-color="primary" indicator-color="primary">
            <!-- TODO: figure out how to make URL query determine the active tab.
            <q-route-tab name="addr" :to="{ query: { mode: 'addr' } }" label="Address" />
            <q-route-tab name="bin" :to="{ query: { mode: 'bin' } }" label="Binary" />
            <q-route-tab name="tx" :to="{ query: { mode: 'tx' } }" label="Tx" />
            -->
            <q-tab default name="addr" label="Address" />
            <q-tab name="bin" label="Binary" />
            <q-tab name="tx" label="Tx" />
          </q-tabs>
          <q-separator />
          <q-tab-panels v-model="mode" animated>

            <q-tab-panel name="addr">
              <div class="row items-center q-gutter-md justify-evenly">
                <div class="col-11 col-sm-3">
                  <q-select v-model="network" :options="networkNames" label="Network" />
                </div>
                <div class="col-grow">
                  <q-input v-model="addr" filled type="search" size="40"
                           placeholder="contract address 0x1234..." />
                </div>
                <div class="col-auto">
                  <q-btn color="primary" label="Decompile contract" @click="btnAddr"/>
                </div>
              </div>
            </q-tab-panel>

            <q-tab-panel name="bin">
              <div class="row items-center q-gutter-md justify-evenly">
                <div class="col-3" hidden>
                  <q-select v-model="format" :options="formatNames" label="Format" />
                </div>
                <div class="col-12">
                  <q-input v-model="bin" filled type="textarea" cols="64"
                           placeholder="60806040..."/>
                </div>
                <div class="col-auto">
                  <q-btn color="primary" label="Decompile contract" @click="btnBin"/>
                </div>
              </div>
            </q-tab-panel>

            <q-tab-panel name="tx">
              Paste data.
            </q-tab-panel>

          </q-tab-panels>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<style>
</style>

<script>

const networks = [
  { label: 'Ethereum mainnet', value: 'eth-mainnet' },
  { label: 'Klaytn Cypress mainnet', value: 'klaytn-mainnet' },
  { label: 'Ethereum Ropsten testnet', value: 'eth-ropsten' },
];

const formats = [
  { label: 'EVM', value: 'evm-generic' },
];

export default {
  name: 'Search',
  data() {
    return {
      mode: 'addr',

      networkNames: networks,
      network: networks[0],
      addr: '',

      formatNames: formats,
      format: formats[0],
      bin: '',
    };
  },
  created() {
  },
  methods: {
    btnAddr: function() {
      var vm = this;
      var path = "/code/addr/" + vm.network.value + "-" + vm.addr;
      vm.$router.push({
        path: path
      });
    },
    btnBin: function() {
      var vm = this;
      console.log(vm.format.value, vm.bin);
      location.href="/#/code/bin/evm-generic-6d967f98f2f3843065688dc2065248e3686b56fc0b6ddfa82007df016148becb"
    },
  },
}
</script>
