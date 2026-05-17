<template>
  <div class="flex h-screen bg-[#f8fafc] font-sans text-slate-800 overflow-hidden">
    
    <!-- Sidebar Layout -->
    <aside class="w-[260px] bg-[#1e293b] text-slate-300 flex flex-col shrink-0">
      <div class="h-16 flex items-center px-6 text-white font-bold text-xl tracking-wide border-b border-slate-700/50">
        <span class="text-orange-400 mr-2 text-2xl">⚡</span> techDash
      </div>
      
      <nav class="flex-1 px-3 py-6 space-y-1 overflow-y-auto">
        <button 
          @click="activeTab = 'dashboard'"
          :class="['w-full flex items-center gap-3 px-3 py-2.5 rounded-lg font-medium transition-colors text-left', activeTab === 'dashboard' ? 'bg-[#334155] text-white' : 'hover:bg-[#334155]/50 hover:text-white']"
        >
          <span class="text-lg">📊</span> Дашборд
        </button>
        <button 
          @click="activeTab = 'organizations'"
          :class="['w-full flex items-center gap-3 px-3 py-2.5 rounded-lg font-medium transition-colors text-left', activeTab === 'organizations' ? 'bg-[#334155] text-white' : 'hover:bg-[#334155]/50 hover:text-white']"
        >
          <span class="text-lg">🏢</span> Организации
        </button>
        <button 
          @click="activeTab = 'services'"
          :class="['w-full flex items-center gap-3 px-3 py-2.5 rounded-lg font-medium transition-colors text-left', activeTab === 'services' ? 'bg-[#334155] text-white' : 'hover:bg-[#334155]/50 hover:text-white']"
        >
          <span class="text-lg">⚙️</span> Справочник услуг
        </button>
        <!-- Вкладка Аналитики -->
        <button 
          @click="activeTab = 'analytics'"
          :class="['w-full flex items-center gap-3 px-3 py-2.5 rounded-lg font-medium transition-colors text-left', activeTab === 'analytics' ? 'bg-[#334155] text-white' : 'hover:bg-[#334155]/50 hover:text-white']"
        >
          <span class="text-lg">📈</span> Аналитика
        </button>
      </nav>
      
      <div class="p-3 border-t border-slate-700/50">
        <button class="w-full flex items-center gap-3 px-3 py-2.5 hover:bg-slate-800 hover:text-white rounded-lg font-medium transition-colors text-left">
          <span class="text-lg">🚪</span> Выход
        </button>
      </div>
    </aside>

    <!-- Main Content Layout -->
    <main class="flex-1 flex flex-col h-screen overflow-hidden bg-white">
      <header class="h-16 bg-white border-b border-slate-200 flex items-center justify-between px-8 shrink-0">
        <h1 class="text-xl font-bold text-slate-900">{{ pageTitle }}</h1>
        
        <div class="flex items-center gap-6">
          <div class="relative w-64 hidden md:block">
            <!-- Two-way binding for search query -->
            <input 
              v-model="searchQuery"
              type="text" 
              placeholder="Быстрый поиск..." 
              class="w-full pl-10 pr-4 py-2 bg-slate-50 border border-slate-200 rounded-full text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white transition-colors"
            />
            <span class="absolute left-3.5 top-2 text-slate-400">🔍</span>
          </div>
          
          <div class="flex items-center gap-3 cursor-pointer">
            <img src="https://i.pravatar.cc/150?img=11" alt="Avatar" class="w-10 h-10 rounded-full border border-slate-200">
            <div class="hidden sm:block">
              <p class="font-semibold text-slate-800 text-sm leading-tight">Admin User</p>
              <p class="text-slate-500 text-xs">Системный администратор</p>
            </div>
          </div>
        </div>
      </header>

      <!-- Dynamic View Injection -->
      <div class="flex-1 overflow-auto p-8 bg-[#f8fafc]">
        <!-- Render the component based on activeTab and pass down the searchQuery as a prop -->
        <DashboardView v-if="activeTab === 'dashboard'" :searchQuery="searchQuery" />
        <OrganizationsView v-if="activeTab === 'organizations'" :searchQuery="searchQuery" />
        <ServicesView v-if="activeTab === 'services'" :searchQuery="searchQuery" />
        <AnalyticsView v-if="activeTab === 'analytics'" :searchQuery="searchQuery" />
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { fetchAllData } from './store.js'

// Import our new view components
import DashboardView from './views/DashboardView.vue'
import OrganizationsView from './views/OrganizationsView.vue'
import ServicesView from './views/ServicesView.vue'
import AnalyticsView from './views/AnalyticsView.vue'

// Global App State
const activeTab = ref('dashboard')
const searchQuery = ref('')

// Load backend data once when app starts
onMounted(() => {
  fetchAllData()
})

// Clear search input when switching tabs
watch(activeTab, () => {
  searchQuery.value = ''
})

// Dynamic title resolution
const pageTitle = computed(() => {
  if (activeTab.value === 'dashboard') return 'Обзор испытаний'
  if (activeTab.value === 'organizations') return 'Справочник организаций'
  if (activeTab.value === 'services') return 'Виды испытаний'
  if (activeTab.value === 'analytics') return 'Аналитика и отчеты'
  return ''
})
</script>