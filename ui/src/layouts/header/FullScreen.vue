<script setup lang="ts">
import { ref } from 'vue'

import { IconArrowsMaximize, IconArrowsMinimize } from '@tabler/icons-vue'

const fullscreen = ref(false)

function toggleFullscreen () {
  fullscreen.value = !fullscreen.value
  if (document && !document.fullscreenElement) {
    document.documentElement.requestFullscreen()
  } else if (document.exitFullscreen) {
    document.exitFullscreen()
  }
}
</script>
<template>
  <div>
    <v-btn
      class="text-secondary ml-sm-2 ml-1"
      color="darkText"
      :icon="fullscreen ? IconArrowsMinimize:IconArrowsMaximize"
      rounded="sm"
      size="small"
      @click="toggleFullscreen"
    />
  </div>
</template>
