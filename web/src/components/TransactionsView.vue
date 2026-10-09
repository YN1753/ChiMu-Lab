<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { Plus, ChevronLeft, ChevronRight, Trash2, Edit2, Wallet } from 'lucide-vue-next'
import { fetchTransactions, fetchTransactionSummary, deleteTransaction, updateTransaction } from '../utils/api'
import type { Transaction, TransactionSummary } from '../types'

const props = defineProps<{
  transactions?: Transaction[]
}>()

const emit = defineEmits<{
  (e: 'openAdd'): void
  (e: 'transactionDeleted', id: number): void
  (e: 'transactionUpdated', tx: Transaction): void
}>()

const currentDate = ref(new Date())
const transactionsList = ref<Transaction[]>([])
const summary = ref<TransactionSummary | null>(null)
const isLoading = ref(false)

// 编辑态
const editingTx = ref<Transaction | null>(null)
const editTitle = ref('')
const editAmount = ref('')
const editCategory = ref('餐饮')
const editType = ref<'expense' | 'income'>('expense')
const editNote = ref('')

const year = computed(() => currentDate.value.getFullYear())
const month = computed(() => currentDate.value.getMonth() + 1)

const prevMonth = () => {
  currentDate.value = new Date(year.value, month.value - 2, 1)
}

const nextMonth = () => {
  currentDate.value = new Date(year.value, month.value, 1)
}

const loadData = async () => {
  try {
    isLoading.value = true
    const [txs, sum] = await Promise.all([
      fetchTransactions({ year: year.value, month: month.value }),
      fetchTransactionSummary({ year: year.value, month: month.value }),
    ])
    transactionsList.value = txs
    summary.value = sum
  } catch (err) {
    console.error('Failed to load transactions:', err)
  } finally {
    isLoading.value = false
  }
}

watch([year, month], () => {
  loadData()
})

onMounted(() => {
  loadData()
})

// 分组：按天汇聚
interface DayGroup {
  dateStr: string
  displayDay: string
  items: Transaction[]
  dayTotalExpense: number
}

const dayGroups = computed(() => {
  const map = new Map<string, Transaction[]>()
  transactionsList.value.forEach((tx) => {
    const d = new Date(tx.occurred_at)
    const key = !isNaN(d.getTime())
      ? `${d.getMonth() + 1}月${d.getDate()}日`
      : '近期'
    const list = map.get(key) || []
    list.push(tx)
    map.set(key, list)
  })

  const groups: DayGroup[] = []
  map.forEach((items, displayDay) => {
    let dayTotal = 0
    items.forEach((it) => {
      if (it.type === 'expense') dayTotal += it.amount
    })
    groups.push({
      dateStr: displayDay,
      displayDay,
      items,
      dayTotalExpense: dayTotal,
    })
  })

  return groups
})

const formatMoney = (cents: number): string => {
  return (cents / 100).toLocaleString('zh-CN', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  })
}

const startEdit = (tx: Transaction) => {
  editingTx.value = tx
  editTitle.value = tx.title
  editAmount.value = (tx.amount / 100).toFixed(2)
  editCategory.value = tx.category
  editType.value = tx.type
  editNote.value = tx.note || ''
}

const saveEdit = async () => {
  if (!editingTx.value) return
  const amtFloat = parseFloat(editAmount.value)
  if (isNaN(amtFloat) || amtFloat <= 0) return

  const updated = await updateTransaction(editingTx.value.id, {
    title: editTitle.value.trim() || editCategory.value,
    amount: Math.round(amtFloat * 100),
    type: editType.value,
    category: editCategory.value,
    note: editNote.value.trim(),
  })

  emit('transactionUpdated', updated)
  editingTx.value = null
  loadData()
}

const removeTx = async (id: number) => {
  if (!confirm('确定删除这笔记录？')) return
  await deleteTransaction(id)
  emit('transactionDeleted', id)
  loadData()
}

// 支出光谱细线条颜色与占比 (Earthy publication film palette)
const CATEGORY_COLORS: Record<string, string> = {
  餐饮: '#a24f2b', // 暖赭石
  交通: '#4a6b82', // 青石灰
  购物: '#8c6d46', // 旧皮褐
  娱乐: '#7a5980', // 暮紫
  住房: '#3f5a4a', // 苔藓绿
  学习: '#5c6773', // 墨色
  医疗: '#b35c5c', // 枯叶红
  其他: '#8a857b', // 米灰
}

const spectrumItems = computed(() => {
  if (!summary.value || !summary.value.total_expense || summary.value.total_expense <= 0) return []
  const breakdown = summary.value.category_breakdown || {}
  const total = summary.value.total_expense
  return Object.entries(breakdown)
    .filter(([_, amt]) => amt > 0)
    .map(([cat, amt]) => {
      const percentage = Math.round((amt / total) * 1000) / 10
      return {
        category: cat,
        amount: amt,
        percentage,
        color: CATEGORY_COLORS[cat] || '#8a857b',
      }
    })
    .sort((a, b) => b.amount - a.amount)
})
</script>

<template>
  <div class="max-w-4xl mx-auto px-5 sm:px-8 py-20 sm:py-28 text-left">
    
    <!-- 标头与月度切换器 -->
    <header class="pb-10 border-b border-[var(--border-subtle)] mb-12 space-y-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
            记账
          </h1>
          <span class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
            // BOOKKEEPING
          </span>
        </div>

        <button
          @click="emit('openAdd')"
          class="inline-flex items-center gap-1.5 px-3 py-1.5 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] rounded-xs hover:opacity-90 transition-opacity cursor-pointer"
        >
          <Plus class="w-3.5 h-3.5" />
          <span>记一笔</span>
        </button>
      </div>

      <!-- 月度指示器 -->
      <div class="flex items-center justify-between pt-2">
        <div class="flex items-center gap-3 font-mono-archive text-sm">
          <button
            @click="prevMonth"
            class="p-1 hover:text-[var(--ink-primary)] text-[var(--ink-muted)] cursor-pointer"
            title="上个月"
          >
            <ChevronLeft class="w-4 h-4" />
          </button>
          <span class="text-base text-[var(--ink-primary)] font-serif-editorial font-medium">
            {{ year }} 年 {{ month }} 月
          </span>
          <button
            @click="nextMonth"
            class="p-1 hover:text-[var(--ink-primary)] text-[var(--ink-muted)] cursor-pointer"
            title="下个月"
          >
            <ChevronRight class="w-4 h-4" />
          </button>
        </div>

        <p class="font-serif-editorial text-xs text-[var(--ink-muted)] hidden sm:block">
          记录日常的柴米油盐与生活花销。
        </p>
      </div>
    </header>

    <!-- 月度核心收支概览 (Editorial 优雅排版，非粗笨 SaaS 卡片) -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-8 pb-12 mb-14 border-b border-[var(--border-subtle)] font-serif-editorial">
      <div class="space-y-1">
        <span class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)] block">// 本月支出</span>
        <div class="text-3xl font-medium tracking-tight text-[var(--ink-primary)]">
          ¥{{ formatMoney(summary?.total_expense || 0) }}
        </div>
      </div>

      <div class="space-y-1">
        <span class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)] block">// 本月收入</span>
        <div class="text-3xl font-medium tracking-tight text-[var(--accent-moss)]">
          ¥{{ formatMoney(summary?.total_income || 0) }}
        </div>
      </div>

      <div class="space-y-1">
        <span class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)] block">// 结余</span>
        <div
          class="text-3xl font-medium tracking-tight"
          :class="(summary?.balance || 0) >= 0 ? 'text-[var(--ink-primary)]' : 'text-red-700'"
        >
          {{ (summary?.balance || 0) < 0 ? '-' : '' }}¥{{ formatMoney(Math.abs(summary?.balance || 0)) }}
        </div>
      </div>
    </div>

    <!-- 支出光谱细线条 (Spending Spectrum Bar · Editorial 典雅色带) -->
    <div
      v-if="summary && summary.total_expense > 0 && spectrumItems.length > 0"
      class="pb-10 mb-14 border-b border-[var(--border-subtle)] space-y-4"
    >
      <div class="flex items-center justify-between text-xs font-mono-archive text-[var(--ink-muted)]">
        <span class="tracking-wider">// 支出光谱分布 · SPECTRUM</span>
        <span>{{ spectrumItems.length }} 个支出类目</span>
      </div>

      <!-- 分段光谱细线 (微小圆角，纯色拼接) -->
      <div class="h-2 w-full rounded-full bg-[var(--bg-subtle)] overflow-hidden flex">
        <div
          v-for="item in spectrumItems"
          :key="item.category"
          :style="{ width: `${item.percentage}%`, backgroundColor: item.color }"
          class="h-full transition-all duration-300 relative group cursor-pointer"
          :title="`${item.category}: ¥${formatMoney(item.amount)} (${item.percentage}%)`"
        ></div>
      </div>

      <!-- 类目图例与占比明细 -->
      <div class="flex flex-wrap items-center gap-x-6 gap-y-2 text-xs font-mono-archive">
        <div
          v-for="item in spectrumItems"
          :key="item.category"
          class="flex items-center gap-2 text-[var(--ink-secondary)]"
        >
          <span class="w-2 h-2 rounded-full shrink-0" :style="{ backgroundColor: item.color }"></span>
          <span>{{ item.category }}</span>
          <span class="text-[var(--ink-muted)]">{{ item.percentage }}%</span>
          <span class="font-mono-archive text-[var(--ink-primary)]">¥{{ formatMoney(item.amount) }}</span>
        </div>
      </div>
    </div>

    <!-- 记账列表 (按日汇聚，Editorial 细线排版) -->
    <div v-if="dayGroups.length === 0" class="py-20 text-center space-y-4">
      <div class="w-12 h-12 rounded-full border border-[var(--border-subtle)] flex items-center justify-center mx-auto text-[var(--ink-muted)]">
        <Wallet class="w-5 h-5" />
      </div>
      <div class="space-y-1 font-serif-editorial">
        <h3 class="text-lg text-[var(--ink-primary)]">这个月还没有记录</h3>
        <p class="text-xs text-[var(--ink-muted)]">
          从一顿具体的饭、一杯咖啡开始，记录日常的收支脚步。
        </p>
      </div>
      <button
        @click="emit('openAdd')"
        class="inline-flex items-center gap-1.5 px-4 py-2 text-xs font-mono-archive border border-[var(--border-subtle)] hover:border-[var(--ink-primary)] rounded-xs cursor-pointer transition-colors"
      >
        <Plus class="w-3.5 h-3.5" />
        <span>添加第一笔账目</span>
      </button>
    </div>

    <div v-else class="space-y-12">
      <section
        v-for="group in dayGroups"
        :key="group.dateStr"
        class="space-y-4"
      >
        <!-- 日期锚点 -->
        <div class="flex items-baseline justify-between border-b border-[var(--border-subtle)] pb-2 text-xs font-mono-archive text-[var(--ink-muted)]">
          <span class="text-sm font-serif-editorial text-[var(--ink-primary)] font-medium">
            {{ group.displayDay }}
          </span>
          <span v-if="group.dayTotalExpense > 0">
            日支出：-¥{{ formatMoney(group.dayTotalExpense) }}
          </span>
        </div>

        <!-- 账目细项 -->
        <div class="divide-y divide-[var(--border-subtle)]/60">
          <div
            v-for="tx in group.items"
            :key="tx.id"
            class="py-3.5 flex items-center justify-between group hover:bg-[var(--bg-subtle)]/40 px-2 rounded-xs transition-colors"
          >
            <div class="space-y-0.5">
              <div class="flex items-center gap-2.5">
                <span class="font-serif-editorial text-base text-[var(--ink-primary)]">
                  {{ tx.title }}
                </span>
                <span class="text-[10px] font-mono-archive tracking-wider text-[var(--ink-muted)] border border-[var(--border-subtle)] px-1.5 py-0.2 rounded-xs">
                  {{ tx.category }}
                </span>
              </div>
              <p v-if="tx.note" class="text-xs font-serif-editorial text-[var(--ink-muted)]">
                {{ tx.note }}
              </p>
            </div>

            <!-- 金额与操作 -->
            <div class="flex items-center gap-4 font-mono-archive text-sm">
              <span
                :class="tx.type === 'expense' ? 'text-[var(--ink-primary)] font-medium' : 'text-[var(--accent-moss)] font-medium'"
              >
                {{ tx.type === 'expense' ? '-' : '+' }}¥{{ formatMoney(tx.amount) }}
              </span>

              <div class="hidden group-hover:flex items-center gap-1.5 text-xs text-[var(--ink-muted)]">
                <button
                  @click="startEdit(tx)"
                  class="p-1 hover:text-[var(--ink-primary)] cursor-pointer"
                  title="修改"
                >
                  <Edit2 class="w-3.5 h-3.5" />
                </button>
                <button
                  @click="removeTx(tx.id)"
                  class="p-1 hover:text-red-700 cursor-pointer"
                  title="删除"
                >
                  <Trash2 class="w-3.5 h-3.5" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>

    <!-- 内联编辑模态 -->
    <div
      v-if="editingTx"
      class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-xs"
      @click.self="editingTx = null"
    >
      <div class="bg-[var(--bg-archive)] border border-[var(--border-subtle)] p-6 rounded-xs w-full max-w-md space-y-4">
        <h3 class="font-serif-editorial text-lg font-medium text-[var(--ink-primary)]">
          修改账目
        </h3>
        <div class="space-y-3">
          <div>
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)] mb-1">事项名称</label>
            <input v-model="editTitle" type="text" class="w-full px-3 py-1.5 text-sm border border-[var(--border-subtle)] rounded-xs outline-none bg-transparent" />
          </div>
          <div>
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)] mb-1">金额 (¥)</label>
            <input v-model="editAmount" type="number" step="0.01" class="w-full px-3 py-1.5 text-sm border border-[var(--border-subtle)] rounded-xs outline-none bg-transparent" />
          </div>
          <div>
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)] mb-1">备注</label>
            <input v-model="editNote" type="text" class="w-full px-3 py-1.5 text-sm border border-[var(--border-subtle)] rounded-xs outline-none bg-transparent" />
          </div>
        </div>
        <div class="flex items-center justify-end gap-3 pt-3 border-t border-[var(--border-subtle)]">
          <button @click="editingTx = null" class="px-3 py-1.5 text-xs font-mono-archive text-[var(--ink-muted)] cursor-pointer">取消</button>
          <button @click="saveEdit" class="px-4 py-1.5 text-xs font-mono-archive bg-[var(--ink-primary)] text-white rounded-xs cursor-pointer">保存</button>
        </div>
      </div>
    </div>

  </div>
</template>
