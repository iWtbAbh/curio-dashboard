<script setup lang="ts">

import { formatBytes } from '@/utils/helpers/formatBytes'
import { StoragePath } from '@/typed-graph'
import { ApolloError } from '@apollo/client'
import { getColorByType } from '@/utils/helpers/storageTypeColor'
import moment from 'moment'

const props = defineProps({
  items: {
    type: Array as () => StoragePath[],
    required: true,
  },
  loading: {
    type: Boolean,
    required: true,
  },
  error: {
    type: Object as () => ApolloError | null,
    default: null,
    required: false,
  },
  search: {
    type: String,
    default: null,
    required: false,
  },
})

const headers = [
  { title: 'Type', key: 'type' },
  { title: 'Storage ID', key: 'storageId', sortable: false },
  { title: 'URLs', key: 'urls', sortable: false },
  { title: 'Weight', key: 'weight' },
  { title: 'Capacity', key: 'capacity' },
  { title: 'Available', key: 'available' },
  { title: 'FS Available', key: 'fsAvailable' },
  { title: 'Reserved', key: 'reserved' },
  { title: 'Used', key: 'used' },
  { title: 'Last Heartbeat', key: 'lastHeartbeat' },
]

</script>

<template>
  <v-data-table-virtual
    fixed-header
    :headers="headers"
    :items="props.items"
    :loading="props.loading"
    :search="props.search"
  >
    <template #item.type="{ value }">
      <v-chip
        :color="getColorByType(value)"
        label
        size="small"
      >
        {{ value }}
      </v-chip>
    </template>
    <template #item.lastHeartbeat="{ value }">
      {{ moment(value).fromNow() }}
    </template>
    <template #item.storageId="{ value }">
      <TruncatedText :text="value" />
    </template>
    <template #item.urls="{ value }">
      <v-chip-group column>
        <TruncatedText
          v-for="url in value.split(',')"
          :key="url"
          :text="url"
          :max-length="50"
        />
      </v-chip-group>
    </template>
    <template #item.capacity="{ value }">
      <div>{{ formatBytes(value).combined }}</div>
    </template>
    <template #item.available="{ value }">
      <div>{{ formatBytes(value).combined }}</div>
    </template>
    <template #item.fsAvailable="{ value }">
      <div>{{ formatBytes(value).combined }}</div>
    </template>
    <template #item.reserved="{ value }">
      <div>{{ formatBytes(value).combined }}</div>
    </template>
    <template #item.used="{ value }">
      <div>{{ formatBytes(value).combined }}</div>
    </template>
  </v-data-table-virtual>
</template>

<style scoped lang="scss">

</style>
