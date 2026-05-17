// src/store.js
import { ref } from 'vue'

// Укажи здесь базовый URL твоего бэкенда 
// (например, порт 8080 или 3000, на котором он крутится локально)
const API_BASE_URL = 'http://localhost:8080/api/v1'

// Реактивные переменные для хранения данных
export const organizationsData = ref([])
export const servicesData = ref([])
export const transactionsData = ref([])

// Состояние загрузки (чтобы показывать спиннер, если нужно)
export const isLoading = ref(true)

// Функция для загрузки всех данных разом
export const fetchAllData = async () => {
  isLoading.value = true
  try {
    // Параллельно запускаем все 3 запроса для скорости
    const [orgsRes, servicesRes, transRes] = await Promise.all([
      fetch(`${API_BASE_URL}/organizations`),
      fetch(`${API_BASE_URL}/services`),
      fetch(`${API_BASE_URL}/acts`)
    ])

    if (!orgsRes.ok || !servicesRes.ok || !transRes.ok) {
      throw new Error('Ошибка при получении данных с сервера')
    }

    // Распаковываем JSON и кладем в наши реактивные переменные
    organizationsData.value = await orgsRes.json()
    servicesData.value = await servicesRes.json()
    transactionsData.value = await transRes.json()
    
  } catch (error) {
    console.error("Сбой загрузки:", error)
    // Тут в будущем можно добавить всплывающее уведомление об ошибке
  } finally {
    isLoading.value = false
  }
}

async function postJson(url, body) {
  const res = await fetch(url, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })

  if (!res.ok) {
    const errorBody = await res.json().catch(() => ({}))
    throw new Error(errorBody.error || 'Ошибка сервера')
  }

  return res.json()
}

export const createOrganization = async (org) => {
  const created = await postJson(`${API_BASE_URL}/organizations`, org)
  await fetchAllData()
  return created
}

export const deleteOrganization = async (id) => {
  const res = await fetch(`${API_BASE_URL}/organizations/${id}`, {
    method: 'DELETE',
  })

  if (!res.ok) {
    const errorBody = await res.json().catch(() => ({}))
    throw new Error(errorBody.error || 'Ошибка при удалении организации')
  }

  await fetchAllData()
}

export const createServiceType = async (service) => {
  const created = await postJson(`${API_BASE_URL}/services`, service)
  await fetchAllData()
  return created
}

export const deleteServiceType = async (id) => {
  const res = await fetch(`${API_BASE_URL}/services/${id}`, {
    method: 'DELETE',
  })

  if (!res.ok) {
    const errorBody = await res.json().catch(() => ({}))
    throw new Error(errorBody.error || 'Ошибка при удалении услуги')
  }

  await fetchAllData()
}

export const createAct = async (act) => {
  const created = await postJson(`${API_BASE_URL}/acts`, act)
  await fetchAllData()
  return created
}

export const exportReport = async () => {
  const res = await fetch(`${API_BASE_URL}/reports/export`, {
    method: 'GET',
  })

  if (!res.ok) {
    const errorBody = await res.json().catch(() => ({}))
    throw new Error(errorBody.error || 'Ошибка при экспорте')
  }

  const blob = await res.blob()
  const fileName = 'report.xlsx'
  const link = document.createElement('a')
  link.href = URL.createObjectURL(blob)
  link.download = fileName
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}
