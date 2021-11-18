<template>
  <q-layout view="hHh lpr lfr">
    <q-header>
      <q-toolbar>
        <q-toolbar-title>
          evmexp
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <Search />
    </q-page-container>

    <q-footer>
      <q-toolbar>
        Powered by <b>panoramix decompiler</b>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<script>
import { ref } from 'vue'
import Search from './components/Search.vue'

export default {
  name: 'LayoutDefault',

  components: {
    Search
  },

  setup () {
    return {
      leftDrawerOpen: ref(false)
    }
  }
}
</script>
