<template>
  <div>
    <div class="mb-8 flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h2 class="text-3xl font-bold text-slate-900 mb-1">Организации</h2>
        <p class="text-slate-600 text-sm">Управление клиентской базой</p>
      </div>
      <button @click="openCreateModal" class="bg-slate-800 hover:bg-slate-900 text-white px-5 py-2.5 rounded-lg font-medium transition-colors flex items-center gap-2 shadow-sm">
        <span class="text-lg leading-none">+</span> Добавить компанию
      </button>
    </div>

    <div v-if="showModal" class="fixed inset-0 z-40 flex items-center justify-center bg-black/40 p-4">
      <div class="w-full max-w-md rounded-3xl bg-white p-6 shadow-2xl">
        <div class="flex items-center justify-between mb-4">
          <div>
            <h3 class="text-xl font-bold text-slate-900">{{ isEdit ? 'Редактировать организацию' : 'Добавить компанию' }}</h3>
            <p class="text-sm text-slate-500">Заполните данные и сохраните.</p>
          </div>
          <button @click="closeModal" class="text-slate-400 hover:text-slate-600">✕</button>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">Наименование</label>
            <input v-model="form.name" type="text" class="w-full rounded-2xl border border-slate-200 px-4 py-3 text-sm focus:border-slate-500 focus:outline-none" />
          </div>
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">БИН</label>
            <input v-model="form.bin" type="text" maxlength="12" class="w-full rounded-2xl border border-slate-200 px-4 py-3 text-sm focus:border-slate-500 focus:outline-none" />
          </div>
          <p v-if="errorMessage" class="text-sm text-red-500">{{ errorMessage }}</p>
        </div>

        <div class="mt-6 flex justify-end gap-3">
          <button @click="closeModal" class="rounded-2xl border border-slate-200 px-4 py-2 text-sm text-slate-700 hover:bg-slate-50">Отмена</button>
          <button @click="saveOrganization" class="rounded-2xl bg-slate-800 px-4 py-2 text-sm font-medium text-white hover:bg-slate-900">Сохранить</button>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-2xl shadow-sm border border-slate-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse whitespace-nowrap">
          <thead>
            <tr class="bg-slate-100 text-slate-600 text-sm border-b border-slate-200">
              <th class="p-4 pl-6 font-semibold w-24">ID</th>
              <th class="p-4 font-semibold">Наименование организации</th>
              <th class="p-4 font-semibold">БИН</th>
              <th class="p-4 font-semibold text-right pr-6">Действия</th>
            </tr>
          </thead>
          <tbody class="text-sm divide-y divide-slate-100 text-slate-700">
            <tr v-for="org in filteredOrganizations" :key="org.id" class="hover:bg-slate-50 transition-colors">
              <td class="p-4 pl-6 font-medium text-slate-500">#{{ org.id }}</td>
              <td class="p-4 font-semibold text-slate-800">{{ org.name }}</td>
              <td class="p-4 font-mono text-slate-600">{{ org.bin }}</td>
              <td class="p-4 pr-6 text-right">
                <button @click="openEditModal(org)" class="text-blue-500 hover:text-blue-700 mr-3">Редактировать</button>
                <button @click="confirmDeleteOrganization(org)" class="text-red-500 hover:text-red-700">Удалить</button>
              </td>
            </tr>
            <tr v-if="filteredOrganizations.length === 0">
              <td colspan="4" class="p-8 text-center text-slate-500">По запросу "{{ searchQuery }}" ничего не найдено.</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { organizationsData, createOrganization, deleteOrganization } from '../store.js'

const showModal = ref(false)
const isEdit = ref(false)
const errorMessage = ref('')
const form = ref({ id: null, name: '', bin: '' })

// Receive search query from parent
const props = defineProps({
  searchQuery: {
    type: String,
    default: ''
  }
})

const openCreateModal = () => {
  form.value = { id: null, name: '', bin: '' }
  errorMessage.value = ''
  isEdit.value = false
  showModal.value = true
}

const openEditModal = (org) => {
  form.value = { id: org.id, name: org.name, bin: org.bin }
  errorMessage.value = ''
  isEdit.value = true
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}

const saveOrganization = async () => {
  if (!form.value.name || form.value.bin.length !== 12) {
    errorMessage.value = 'Введите корректное название и 12-значный БИН.'
    return
  }

  try {
    await createOrganization({ name: form.value.name, bin: form.value.bin })
    closeModal()
  } catch (error) {
    errorMessage.value = error.message || 'Ошибка при сохранении организации.'
  }
}

const confirmDeleteOrganization = async (org) => {
  if (!window.confirm(`Удалить организацию «${org.name}»?`)) {
    return
  }

  try {
    await deleteOrganization(org.id)
  } catch (error) {
    alert(error.message || 'Ошибка при удалении организации.')
  }
}

// Filter organizations based on search query
const filteredOrganizations = computed(() => {
  if (!props.searchQuery) return organizationsData.value
  const query = props.searchQuery.toLowerCase()
  return organizationsData.value.filter(org => {
    return org.name.toLowerCase().includes(query) || org.bin.includes(query)
  })
})
</script>