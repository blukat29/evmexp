<template>
  <q-page class="flex flex-center">
    <div class="row">
      <div class="col-3">
        <q-select v-model="network" :options="networkNames" label="Network" />
      </div>
      <div class="col-9">
        <q-input v-model="search" filled type="search" size="64"
                 placeholder="contract address, txid, 0x1234..." />
      </div>
      <div class="col-12">
        <q-list bordered class="rounded-borders">
          <q-expansion-item dense dense-toggle expand-separator label="Manual input">
            <q-input v-model="manualInput" filled type="textarea">
            </q-input>
          </q-expansion-item>
        </q-list>
      </div>
    </div>
    <div class="row justify-end">
      <div class="col-3">
        <q-btn color="primary" label="Go" @click="btnGo"/>
      </div>
    </div>
  </q-page>
</template>

<style>
</style>

<script>
export default {
  name: 'Search',
  data() {
    return {
      network: "",
      search: "",
      isManual: false,
      manualInput: "",
      networkNames: ["Klaytn", "Ethereum"],
    };
  },
  methods: {
    btnGo: function() {
      var vm = this;
      console.log(vm.isManual, vm.network, vm.search, vm.manualInput);
      location.href="/#/contract"
    },
  },
}
</script>
