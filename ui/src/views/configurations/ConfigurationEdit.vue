<script setup lang="ts">

import { computed, ComputedRef, ref, watch } from 'vue'
import { useMutation, useQuery } from '@vue/apollo-composable'
import { GetConfig, UpdateConfig } from '@/gql/config'
import { Config } from '@/typed-graph'
import { useI18n } from 'vue-i18n'
import { IconDeviceFloppy,  IconPencil, IconSettingsCheck } from '@tabler/icons-vue'
import { useNotificationStore } from '@/stores/notification'

const notificationStore = useNotificationStore()
const { t } = useI18n()

const props = defineProps({
  layer: {
    type: String,
    required: true,
  },
})

const enableEdit = ref(false)

const { result, loading, error } = useQuery(GetConfig, {
  layer: props.layer,
}, () => ({
  fetchPolicy: 'cache-first',
}))

const config: ComputedRef<Config> = computed(() => result.value?.config)

const editConfig = ref(config.value?.config)
const editTitle = ref(config.value?.title)
const isValid = ref(false)

watch(result, () => {
  editConfig.value = config.value?.config
  editTitle.value = config.value?.title
})

const { mutate: updateConfig, loading: updateLoading, onDone } = useMutation(UpdateConfig)

function saveEdit () {
  updateConfig({
    title: editTitle.value,
    config: editConfig.value,
  })
}

onDone(() => {
  enableEdit.value = false
  notificationStore.success(`Configuration updated successfully`)
})

const breadcrumbs = ref([
  {
    title: t('nav.Configurations'),
    disabled: false,
    to: { name: 'Configurations' }
  }
])

function handleIsValid(v: boolean) {
  console.log('Config is valid:', v)
  isValid.value = v
}

</script>

<template>
  <BaseBreadcrumb :breadcrumbs="breadcrumbs" />
  <UiParentCard
    :loading="loading"
    :title="t('fields.Edit Configuration')"
  >
    <template #action>
      <v-btn
        v-if="enableEdit"
        color="primary"
        :loading="updateLoading"
        :disabled="!isValid"
        @click="saveEdit"
      >
        <template #prepend>
          <IconDeviceFloppy />
        </template>
        {{ t('actions.Save') }}
      </v-btn>
      <v-btn
        v-else
        color="error"
        @click="enableEdit = true"
      >
        <template #prepend>
          <IconPencil />
        </template>
        {{ t('actions.Edit') }}
      </v-btn>
    </template>
    <v-label class="mb-1">
      {{ t('fields.Layer') }}
    </v-label>
    <v-text-field
      v-model="editTitle"
      color="primary"
      :disabled="true"
      :error-messages="error?.message"
      persistent-hint
      persistent-placeholder
      placeholder="Enter config layer"
      required
      variant="outlined"
    />
    <v-label class="mb-1 mt-5 mr-5">
      {{ t('fields.Content') }}
    </v-label>
    <a
      href="https://docs.curiostorage.org/configuration/default-curio-configuration"
      target="_blank"
      class="text-decoration-none"
    >
      <v-icon
        :icon="IconSettingsCheck"
        size="small"
      />
      Check the default configuration
    </a>
    <ConfigEditor
      v-model="editConfig"
      :readonly="!enableEdit"
      @is-valid="handleIsValid"
    />
  </uiparentcard>
</template>

<style scoped lang="scss">

</style>
