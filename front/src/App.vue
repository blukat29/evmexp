<template>
  <q-layout view="hHh lpr lfr">
    <q-header>
      <q-toolbar class="bg-grey-3 text-dark">
        <q-toolbar-title>
          <router-link style="text-decoration: none; color: inherit;" to="/">Que contract</router-link>
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view></router-view>
    </q-page-container>

    <q-footer>
      <q-toolbar class="bg-grey-3 text-dark">
        Powered by <b>panoramix decompiler</b>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<script>
import { ref } from 'vue'

export default {
  name: 'LayoutDefault',

  components: {
  },

  setup () {
    return {
      leftDrawerOpen: ref(false)
    }
  }
}
</script>
