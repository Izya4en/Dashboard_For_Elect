<template>
  <div class="flex flex-col gap-6">

    <div class="flex flex-col sm:flex-row sm:items-end justify-between gap-4">
      <div>
        <h2 class="text-3xl font-bold text-slate-900 mb-1">Аналитический центр</h2>
        <p class="text-slate-600 text-sm">Агрегация данных и построение отчетов</p>
      </div>
      <button @click="fetchData" class="bg-slate-800 hover:bg-slate-900 text-white px-6 py-2.5 rounded-lg font-medium transition-colors text-sm flex items-center gap-2 self-start sm:self-auto">
        <span>📥</span> Обновить данные
      </button>
    </div>

    <div v-if="isLoading" class="flex justify-center items-center py-24">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
    </div>

    <template v-else>
      <!-- KPI Cards -->
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-4">
        <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-5">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider mb-2">Всего транзакций</p>
          <p class="text-4xl font-bold text-slate-900">{{ totalActs }}</p>
        </div>
        <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-5">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider mb-2">Активных организаций</p>
          <p class="text-4xl font-bold text-blue-600">{{ activeOrgs }}</p>
          <p class="text-xs text-slate-400 mt-1">из {{ totalOrgs }} в справочнике</p>
        </div>
        <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-5">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider mb-2">Видов услуг</p>
          <p class="text-4xl font-bold text-indigo-600">{{ totalServiceTypes }}</p>
        </div>
        <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-5">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider mb-2">Период</p>
          <p class="text-xl font-bold text-slate-700 leading-tight">{{ period }}</p>
        </div>
      </div>

      <!-- Line chart: temporal dynamics -->
      <div class="bg-white p-6 rounded-2xl shadow-sm border border-slate-200">
        <h3 class="font-bold text-slate-800 mb-1">Динамика транзакций по месяцам</h3>
        <p class="text-xs text-slate-400 mb-4">Количество актов испытаний в разрезе месяцев</p>
        <apexchart type="area" height="260" :options="lineOptions" :series="lineSeries" />
      </div>

      <!-- Services + Orgs -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="bg-white p-6 rounded-2xl shadow-sm border border-slate-200">
          <h3 class="font-bold text-slate-800 mb-1">Структура услуг</h3>
          <p class="text-xs text-slate-400 mb-4">Топ-12 по частоте обращений</p>
          <apexchart type="bar" height="380" :options="servicesBarOptions" :series="servicesBarSeries" />
        </div>
        <div class="bg-white p-6 rounded-2xl shadow-sm border border-slate-200">
          <h3 class="font-bold text-slate-800 mb-1">Топ-15 организаций</h3>
          <p class="text-xs text-slate-400 mb-4">По числу транзакций</p>
          <apexchart type="bar" height="380" :options="orgsBarOptions" :series="orgsBarSeries" />
        </div>
      </div>
    </template>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const MONTHS_RU = ['Янв','Фев','Мар','Апр','Май','Июн','Июл','Авг','Сен','Окт','Ноя','Дек']
const fmtMonth = (yyyymm) => {
  const [y, m] = yyyymm.split('-')
  return `${MONTHS_RU[parseInt(m) - 1]} ${y}`
}

const isLoading = ref(true)
const totalActs = ref(0)
const activeOrgs = ref(0)
const totalOrgs = ref(0)
const totalServiceTypes = ref(0)
const period = ref('—')

const lineSeries = ref([{ name: 'Транзакции', data: [] }])
const lineOptions = ref({
  chart: { type: 'area', toolbar: { show: false }, zoom: { enabled: false } },
  dataLabels: { enabled: true, style: { fontSize: '12px', fontWeight: 700, colors: ['#1e40af'] }, background: { enabled: false } },
  stroke: { curve: 'straight', width: 2 },
  fill: { type: 'gradient', gradient: { shadeIntensity: 1, opacityFrom: 0.25, opacityTo: 0.02, stops: [0, 90, 100] } },
  xaxis: { categories: [] },
  yaxis: { labels: { formatter: v => Math.round(v) }, min: 0 },
  markers: { size: 5, colors: ['#3b82f6'], strokeColors: '#fff', strokeWidth: 2 },
  colors: ['#3b82f6'],
  grid: { borderColor: '#f1f5f9' },
  tooltip: { theme: 'light' }
})

const servicesBarSeries = ref([{ name: 'Кол-во', data: [] }])
const servicesBarOptions = ref({
  chart: { type: 'bar', toolbar: { show: false } },
  plotOptions: { bar: { horizontal: true, barHeight: '65%', borderRadius: 3 } },
  dataLabels: { enabled: true, style: { fontSize: '11px' } },
  xaxis: { categories: [] },
  colors: ['#6366f1'],
  grid: { borderColor: '#f1f5f9' },
  tooltip: { theme: 'light' }
})

const orgsBarSeries = ref([{ name: 'Транзакции', data: [] }])
const orgsBarOptions = ref({
  chart: { type: 'bar', toolbar: { show: false } },
  plotOptions: { bar: { borderRadius: 3, columnWidth: '60%' } },
  dataLabels: { enabled: false },
  xaxis: {
    categories: [],
    labels: { rotate: -40, trim: true, maxHeight: 90, style: { fontSize: '11px' } }
  },
  colors: ['#10b981'],
  grid: { borderColor: '#f1f5f9' },
  tooltip: { theme: 'light' }
})

const fetchData = async () => {
  isLoading.value = true
  try {
    const [actsRes, orgsRes] = await Promise.all([
      fetch('http://localhost:8080/api/v1/acts'),
      fetch('http://localhost:8080/api/v1/organizations')
    ])
    const acts = await actsRes.json()
    const orgs = await orgsRes.json()

    totalActs.value = acts.length
    totalOrgs.value = orgs.length
    activeOrgs.value = new Set(acts.map(a => a.bin).filter(Boolean)).size
    totalServiceTypes.value = new Set(acts.map(a => a.service).filter(Boolean)).size

    // --- Line chart ---
    const monthCount = {}
    acts.forEach(act => {
      const m = (act.date || '').substring(0, 7)
      if (m.length === 7) monthCount[m] = (monthCount[m] || 0) + 1
    })

    // Find the dominant year (year with most transactions) to filter out bad dates like 2026-05
    const yearTotals = {}
    Object.entries(monthCount).forEach(([m, c]) => {
      const y = m.substring(0, 4)
      yearTotals[y] = (yearTotals[y] || 0) + c
    })
    const dominantYear = Object.entries(yearTotals).sort((a, b) => b[1] - a[1])[0]?.[0]
    const sortedMonths = Object.keys(monthCount)
      .filter(m => m.startsWith(dominantYear))
      .sort()

    if (sortedMonths.length) {
      const first = fmtMonth(sortedMonths[0])
      const last = fmtMonth(sortedMonths[sortedMonths.length - 1])
      period.value = first === last ? first : `${first} – ${last}`
    }
    lineOptions.value = { ...lineOptions.value, xaxis: { categories: sortedMonths.map(fmtMonth) } }
    lineSeries.value = [{ name: 'Транзакции', data: sortedMonths.map(m => monthCount[m]) }]

    // --- Services horizontal bar (top 12) ---
    const svcCount = {}
    acts.forEach(act => {
      const name = act.service || 'Неизвестно'
      svcCount[name] = (svcCount[name] || 0) + 1
    })
    const topServices = Object.entries(svcCount).sort((a, b) => b[1] - a[1]).slice(0, 12)
    servicesBarOptions.value = { ...servicesBarOptions.value, xaxis: { categories: topServices.map(([n]) => n) } }
    servicesBarSeries.value = [{ name: 'Кол-во', data: topServices.map(([, c]) => c) }]

    // --- Orgs bar (top 15) ---
    const orgCount = {}
    acts.forEach(act => {
      const name = act.organization || 'Неизвестно'
      orgCount[name] = (orgCount[name] || 0) + 1
    })
    const topOrgs = Object.entries(orgCount).sort((a, b) => b[1] - a[1]).slice(0, 15)
    orgsBarOptions.value = {
      ...orgsBarOptions.value,
      xaxis: { ...orgsBarOptions.value.xaxis, categories: topOrgs.map(([n]) => n) }
    }
    orgsBarSeries.value = [{ name: 'Транзакции', data: topOrgs.map(([, c]) => c) }]

  } catch (error) {
    console.error('Analytics fetch error:', error)
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchData)
</script>
