<template>
  <div>
    <div class="mb-8 flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h2 class="text-3xl font-bold text-slate-900 mb-1">Справочник услуг</h2>
        <p class="text-slate-600 text-sm">Перечень проводимых испытаний и проверок</p>
      </div>
      <button @click="openCreateModal" class="bg-emerald-600 hover:bg-emerald-700 text-white px-5 py-2.5 rounded-lg font-medium transition-colors flex items-center gap-2 shadow-sm">
        <span class="text-lg leading-none">+</span> Новая услуга
      </button>
    </div>

    <div v-if="showModal" class="fixed inset-0 z-40 flex items-center justify-center bg-black/40 p-4">
      <div class="w-full max-w-md rounded-3xl bg-white p-6 shadow-2xl">
        <div class="flex items-center justify-between mb-4">
          <div>
            <h3 class="text-xl font-bold text-slate-900">{{ isEdit ? 'Редактировать услугу' : 'Новая услуга' }}</h3>
            <p class="text-sm text-slate-500">Заполните название и сохраните.</p>
          </div>
          <button @click="closeModal" class="text-slate-400 hover:text-slate-600">✕</button>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">Наименование услуги</label>
            <input v-model="form.name" type="text" class="w-full rounded-2xl border border-slate-200 px-4 py-3 text-sm focus:border-slate-500 focus:outline-none" />
          </div>
          <p v-if="errorMessage" class="text-sm text-red-500">{{ errorMessage }}</p>
        </div>

        <div class="mt-6 flex justify-end gap-3">
          <button @click="closeModal" class="rounded-2xl border border-slate-200 px-4 py-2 text-sm text-slate-700 hover:bg-slate-50">Отмена</button>
          <button @click="saveService" class="rounded-2xl bg-emerald-600 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-700">Сохранить</button>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-2xl shadow-sm border border-slate-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse whitespace-nowrap">
          <thead>
            <tr class="bg-slate-100 text-slate-600 text-sm border-b border-slate-200">
              <th class="p-4 pl-6 font-semibold w-24">Код</th>
              <th class="p-4 font-semibold">Наименование услуги</th>
              <th class="p-4 font-semibold text-right pr-6">Действия</th>
            </tr>
          </thead>
          <tbody class="text-sm divide-y divide-slate-100 text-slate-700">
            <tr v-for="service in filteredServices" :key="service.id" class="hover:bg-slate-50 transition-colors">
              <td class="p-4 pl-6 font-medium text-slate-500">#{{ service.id }}</td>
              <td class="p-4 font-medium text-slate-800 whitespace-normal">{{ service.name }}</td>
              <td class="p-4 pr-6 text-right">
                <button @click="openEditModal(service)" class="text-blue-500 hover:text-blue-700 mr-3">Редактировать</button>
                <button @click="confirmDeleteService(service)" class="text-red-500 hover:text-red-700">Удалить</button>
              </td>
            </tr>
            <tr v-if="filteredServices.length === 0">
              <td colspan="3" class="p-8 text-center text-slate-500">По запросу "{{ searchQuery }}" ничего не найдено.</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { servicesData, createServiceType, deleteServiceType } from '../store.js'

// Receive search query from parent
const props = defineProps({
  searchQuery: {
    type: String,
    default: ''
  }
})

const showModal = ref(false)
const isEdit = ref(false)
const errorMessage = ref('')
const form = ref({ id: null, name: '' })

const openCreateModal = () => {
  form.value = { id: null, name: '' }
  errorMessage.value = ''
  isEdit.value = false
  showModal.value = true
}

const openEditModal = (service) => {
  form.value = { id: service.id, name: service.name }
  errorMessage.value = ''
  isEdit.value = true
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}

const saveService = async () => {
  if (!form.value.name) {
    errorMessage.value = 'Введите название услуги.'
    return
  }

  try {
    await createServiceType({ name: form.value.name })
    closeModal()
  } catch (error) {
    errorMessage.value = error.message || 'Ошибка при сохранении услуги.'
  }
}

const confirmDeleteService = async (service) => {
  if (!window.confirm(`Удалить услугу «${service.name}»?`)) {
    return
  }

  try {
    await deleteServiceType(service.id)
  } catch (error) {
    alert(error.message || 'Ошибка при удалении услуги.')
  }
}

// Filter services based on search query
const filteredServices = computed(() => {
  if (!props.searchQuery) return servicesData.value
  const query = props.searchQuery.toLowerCase()
  return servicesData.value.filter(service => {
    return service.name.toLowerCase().includes(query) || String(service.id).includes(query)
  })
})
</script>