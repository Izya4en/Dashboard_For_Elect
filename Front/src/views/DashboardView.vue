<template>
  <div>
    <!-- Welcome Title & Action -->
    <div class="mb-8 flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h2 class="text-3xl font-bold text-slate-900 mb-1">Добро пожаловать в систему</h2>
        <p class="text-slate-600 text-sm">Сводка по техническим испытаниям и проверкам</p>
      </div>
      <div class="flex flex-wrap gap-3">
        <button @click="exportReportFile" class="bg-slate-700 hover:bg-slate-800 text-white px-5 py-2.5 rounded-lg font-medium transition-colors flex items-center gap-2 shadow-sm">
          📥 Экспорт отчета
        </button>
        <button @click="openCreateActModal" class="bg-blue-500 hover:bg-blue-600 text-white px-5 py-2.5 rounded-lg font-medium transition-colors flex items-center gap-2 shadow-sm">
          <span class="text-lg leading-none">+</span> Создать акт
        </button>
      </div>
    </div>

    <div v-if="showActModal" class="fixed inset-0 z-40 flex items-center justify-center bg-black/40 p-4">
      <div class="w-full max-w-xl rounded-3xl bg-white p-6 shadow-2xl">
        <div class="flex items-center justify-between mb-4">
          <div>
            <h3 class="text-xl font-bold text-slate-900">Создать новый акт</h3>
            <p class="text-sm text-slate-500">Укажите дату, организацию и вид услуги.</p>
          </div>
          <button @click="closeCreateActModal" class="text-slate-400 hover:text-slate-600">✕</button>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">Дата</label>
            <input v-model="form.test_date" type="date" class="w-full rounded-2xl border border-slate-200 px-4 py-3 text-sm focus:border-slate-500 focus:outline-none" />
          </div>
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">Организация</label>
            <select v-model.number="form.organization_id" class="w-full rounded-2xl border border-slate-200 px-4 py-3 text-sm focus:border-slate-500 focus:outline-none">
              <option value="" disabled>Выберите организацию</option>
              <option v-for="org in organizationsData" :key="org.id" :value="org.id">{{ org.name }} — {{ org.bin }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">Вид услуги</label>
            <select v-model.number="form.service_type_id" class="w-full rounded-2xl border border-slate-200 px-4 py-3 text-sm focus:border-slate-500 focus:outline-none">
              <option value="" disabled>Выберите услугу</option>
              <option v-for="service in servicesData" :key="service.id" :value="service.id">{{ service.name }}</option>
            </select>
          </div>
          <p v-if="errorMessage" class="text-sm text-red-500">{{ errorMessage }}</p>
        </div>

        <div class="mt-6 flex justify-end gap-3">
          <button @click="closeCreateActModal" class="rounded-2xl border border-slate-200 px-4 py-2 text-sm text-slate-700 hover:bg-slate-50">Отмена</button>
          <button @click="saveAct" class="rounded-2xl bg-blue-500 px-4 py-2 text-sm font-medium text-white hover:bg-blue-600">Сохранить</button>
        </div>
      </div>
    </div>

    <!-- KPI Cards Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="bg-[#eff6ff] rounded-2xl p-6 border border-[#bfdbfe] shadow-sm flex flex-col justify-between">
        <div class="flex items-start gap-4 mb-4">
          <div class="w-12 h-12 rounded-xl bg-[#dbeafe] text-[#3b82f6] flex items-center justify-center text-2xl">🏢</div>
          <div class="text-slate-700 font-medium leading-snug">Клиенты<br>(Организации)</div>
        </div>
        <div class="text-4xl font-bold text-slate-900">{{ organizationsData.length }}</div>
      </div>
      
      <div class="bg-[#ecfdf5] rounded-2xl p-6 border border-[#a7f3d0] shadow-sm flex flex-col justify-between">
        <div class="flex items-start gap-4 mb-4">
          <div class="w-12 h-12 rounded-xl bg-[#d1fae5] text-[#10b981] flex items-center justify-center text-2xl">☑️</div>
          <div class="text-slate-700 font-medium leading-snug">Оказано услуг<br>&nbsp;</div>
        </div>
        <div class="text-4xl font-bold text-slate-900">{{ transactionsData.length }}</div>
      </div>
      
      <div class="bg-[#fffbeb] rounded-2xl p-6 border border-[#fde68a] shadow-sm flex flex-col justify-between">
        <div class="flex items-start gap-4 mb-4">
          <div class="w-12 h-12 rounded-xl bg-[#fef3c7] text-[#f59e0b] flex items-center justify-center text-2xl">⚡</div>
          <div class="text-slate-700 font-medium leading-snug">Виды испытаний<br>&nbsp;</div>
        </div>
        <div class="text-4xl font-bold text-slate-900">{{ servicesData.length }}</div>
      </div>
    </div>

    <!-- Data Table -->
    <div class="bg-white rounded-2xl shadow-sm border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100">
        <h3 class="font-bold text-lg text-slate-800">Последние транзакции</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse whitespace-nowrap">
          <thead>
            <tr class="bg-slate-100 text-slate-600 text-sm border-b border-slate-200">
              <th class="p-4 pl-6 font-semibold">ID</th>
              <th class="p-4 font-semibold">Дата</th>
              <th class="p-4 font-semibold">Организация</th>
              <th class="p-4 font-semibold">БИН</th>
              <th class="p-4 font-semibold">Услуга</th>
            </tr>
          </thead>
          <tbody class="text-sm divide-y divide-slate-100 text-slate-700">
            <tr v-for="tr in filteredTransactions" :key="tr.id" class="hover:bg-slate-50 transition-colors">
              <td class="p-4 pl-6 font-medium text-slate-500">#{{ tr.id }}</td>
              <td class="p-4">{{ tr.date }}</td>
              <td class="p-4 font-medium">{{ tr.organization }}</td>
              <td class="p-4">{{ tr.bin }}</td>
              <td class="p-4 truncate max-w-md" :title="tr.service">
                {{ tr.service }}
              </td>
            </tr>
            <tr v-if="filteredTransactions.length === 0">
              <td colspan="5" class="p-8 text-center text-slate-500">По запросу ничего не найдено.</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { organizationsData, servicesData, transactionsData, createAct, exportReport } from '../store.js'

// Accept search query from parent (App.vue)
const props = defineProps({
  searchQuery: {
    type: String,
    default: ''
  }
})

const showActModal = ref(false)
const errorMessage = ref('')
const form = ref({ test_date: '', organization_id: '', service_type_id: '' })

const openCreateActModal = () => {
  form.value = { test_date: '', organization_id: '', service_type_id: '' }
  errorMessage.value = ''
  showActModal.value = true
}

const closeCreateActModal = () => {
  showActModal.value = false
}

const saveAct = async () => {
  if (!form.value.test_date || !form.value.organization_id || !form.value.service_type_id) {
    errorMessage.value = 'Заполните все поля формы.'
    return
  }

  try {
    await createAct({
      test_date: form.value.test_date,
      organization_id: form.value.organization_id,
      service_type_id: form.value.service_type_id,
    })
    closeCreateActModal()
  } catch (error) {
    errorMessage.value = error.message || 'Ошибка при создании акта.'
  }
}

const exportReportFile = async () => {
  try {
    await exportReport()
  } catch (error) {
    console.error('Ошибка экспорта:', error)
    alert(error.message || 'Не удалось скачать отчет.')
  }
}

// Filtering logic
const filteredTransactions = computed(() => {
  if (!props.searchQuery) return transactionsData.value
  const query = props.searchQuery.toLowerCase()
  return transactionsData.value.filter(t => {
    return t.organization?.toLowerCase().includes(query) ||
           t.bin?.includes(query) ||
           t.service?.toLowerCase().includes(query)
  })
})
</script>