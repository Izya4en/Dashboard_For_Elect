export const organizationsData = [
  { id: 1, name: 'ООО ТехноЛаб', bin: '060840000123' },
  { id: 2, name: 'ЗАО ЭлектроСервис', bin: '060840000456' },
  { id: 3, name: 'ИП Петров П.П.', bin: '060840000789' }
]

export const servicesData = [
  { id: 1, name: 'Проверка электробезопасности' },
  { id: 2, name: 'Испытание трансформаторов' },
  { id: 3, name: 'ТО оборудования' }
]

export const transactionsData = [
  {
    transaction_id: 101,
    date: '2026-05-08',
    organization_id: 1,
    service_id: 2
  },
  {
    transaction_id: 102,
    date: '2026-05-09',
    organization_id: 3,
    service_id: 1
  },
  {
    transaction_id: 103,
    date: '2026-05-10',
    organization_id: 2,
    service_id: 3
  }
]
