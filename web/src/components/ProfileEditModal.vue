<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { X, Check, Trash2, Plus, RotateCcw } from 'lucide-vue-next'
import type { UserProfile, GearItem } from '../types'
import { DEFAULT_PROFILE, saveCuratorProfile } from '../utils/profile'

const props = defineProps<{
  isOpen: boolean
  profile: UserProfile
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'updated', profile: UserProfile): void
}>()

const form = ref<UserProfile>({ ...props.profile })
const activeTab = ref<'base' | 'gear'>('base')

// 装备新增区
const newGearCategory = ref('影像与暗房')
const newGearName = ref('')
const newGearSpec = ref('')
const newGearStatus = ref('日常自用')
const newGearNote = ref('')
const showAddGear = ref(false)

const categories = [
  '影像与暗房',
  '工作台与算力',
  '手冲咖啡工坊',
  '夜骑巡航',
  '纸笔与日常',
]

watch(
  () => props.profile,
  (newP) => {
    form.value = JSON.parse(JSON.stringify(newP))
  },
  { deep: true, immediate: true }
)

const handleSave = () => {
  saveCuratorProfile(form.value)
  emit('updated', form.value)
  emit('close')
}

const handleResetDefault = () => {
  if (!confirm('确定恢复为默认主理人资料与预置设备清单？')) return
  form.value = JSON.parse(JSON.stringify(DEFAULT_PROFILE))
}

const handleAddGear = () => {
  if (!newGearName.value.trim()) return
  const item: GearItem = {
    id: `gear-${Date.now()}`,
    category: newGearCategory.value,
    name: newGearName.value.trim(),
    spec: newGearSpec.value.trim(),
    status: newGearStatus.value.trim(),
    note: newGearNote.value.trim(),
  }
  form.value.gearList.push(item)
  newGearName.value = ''
  newGearSpec.value = ''
  newGearNote.value = ''
  showAddGear.value = false
}

const removeGear = (idx: number) => {
  form.value.gearList.splice(idx, 1)
}

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Escape' && props.isOpen) {
    emit('close')
  }
}

onMounted(() => {
  window.addEventListener('keydown', handleKeydown)
})
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6 bg-black/45 backdrop-blur-xs animate-in fade-in duration-200"
    @click.self="emit('close')"
  >
    <div
      class="bg-[var(--bg-archive)] border border-[var(--border-subtle)] shadow-[0_20px_60px_rgba(0,0,0,0.15)] w-full max-w-2xl max-h-[92vh] overflow-y-auto rounded-xs p-6 sm:p-8 text-left space-y-6"
    >
      <!-- 弹窗顶栏 -->
      <div class="flex items-center justify-between border-b border-[var(--border-subtle)] pb-4">
        <div>
          <h2 class="font-serif-editorial text-2xl font-medium text-[var(--ink-primary)]">
            编辑主理人档案与设备
          </h2>
          <span class="font-mono-archive text-[11px] text-[var(--ink-muted)]">
            // PROFILE & GEAR MANAGER
          </span>
        </div>

        <button
          @click="emit('close')"
          class="p-1 text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
        >
          <X class="w-5 h-5" />
        </button>
      </div>

      <!-- 分页切换：基础资料 · 设备清单 -->
      <div class="flex items-center gap-6 text-xs font-mono-archive border-b border-[var(--border-subtle)] pb-2">
        <button
          @click="activeTab = 'base'"
          class="py-1 cursor-pointer transition-colors"
          :class="activeTab === 'base' ? 'text-[var(--ink-primary)] font-semibold border-b-2 border-[var(--ink-primary)]' : 'text-[var(--ink-muted)] hover:text-[var(--ink-primary)]'"
        >
          主理人基本资料
        </button>
        <button
          @click="activeTab = 'gear'"
          class="py-1 cursor-pointer transition-colors"
          :class="activeTab === 'gear' ? 'text-[var(--ink-primary)] font-semibold border-b-2 border-[var(--ink-primary)]' : 'text-[var(--ink-muted)] hover:text-[var(--ink-primary)]'"
        >
          随身与桌面设备 ({{ form.gearList.length }})
        </button>
      </div>

      <!-- 1. 基础资料表单 -->
      <div v-if="activeTab === 'base'" class="space-y-5">
        <!-- 头像与预览 -->
        <div class="space-y-2">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">
            头像图片链接 (AVATAR URL)
          </label>
          <div class="flex items-center gap-4">
            <div class="w-16 h-16 rounded-xs bg-[var(--bg-subtle)] border border-[var(--border-subtle)] overflow-hidden shrink-0">
              <img
                :src="form.avatar"
                alt="头像预览"
                class="w-full h-full object-cover"
                onerror="this.src='https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=400&q=80'"
              />
            </div>
            <input
              v-model="form.avatar"
              type="text"
              placeholder="https://... 头像网络地址"
              class="flex-grow px-3 py-2 text-xs font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">称谓 / 代号</label>
            <input
              v-model="form.name"
              type="text"
              class="w-full px-3 py-2 text-sm font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>

          <div class="space-y-1">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">身处坐标</label>
            <input
              v-model="form.location"
              type="text"
              class="w-full px-3 py-2 text-sm font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>
        </div>

        <div class="space-y-1">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">角色签名与手艺标签</label>
          <input
            v-model="form.title"
            type="text"
            class="w-full px-3 py-2 text-sm font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
          />
        </div>

        <div class="space-y-1">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">精神信条 (PHILOSOPHY)</label>
          <input
            v-model="form.philosophy"
            type="text"
            class="w-full px-3 py-2 text-sm font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
          />
        </div>

        <div class="space-y-1">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">关于我 · 独白正文</label>
          <textarea
            v-model="form.bio"
            rows="5"
            class="w-full px-3 py-2 text-sm font-serif-editorial leading-relaxed bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none resize-y"
          ></textarea>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">公开联络邮箱</label>
            <input
              v-model="form.email"
              type="text"
              class="w-full px-3 py-2 text-xs font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>

          <div class="space-y-1">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">GitHub 主页链接</label>
            <input
              v-model="form.github"
              type="text"
              class="w-full px-3 py-2 text-xs font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>
        </div>
      </div>

      <!-- 2. 设备清单表单 -->
      <div v-else-if="activeTab === 'gear'" class="space-y-6">
        <!-- 顶部操作栏 -->
        <div class="flex items-center justify-between">
          <span class="text-xs font-mono-archive text-[var(--ink-muted)]">
            已收录 {{ form.gearList.length }} 项设备器物
          </span>

          <button
            @click="showAddGear = !showAddGear"
            type="button"
            class="inline-flex items-center gap-1.5 px-3 py-1 text-xs font-mono-archive border border-[var(--border-subtle)] hover:border-[var(--ink-primary)] rounded-xs cursor-pointer transition-colors"
          >
            <Plus class="w-3.5 h-3.5" />
            <span>{{ showAddGear ? '取消添加' : '＋ 添加新设备' }}</span>
          </button>
        </div>

        <!-- 展开的添加设备小抽屉 -->
        <div
          v-if="showAddGear"
          class="p-4 rounded-xs border border-[var(--border-subtle)] bg-[var(--bg-subtle)]/40 space-y-3"
        >
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
            <div>
              <label class="block text-[11px] font-mono-archive text-[var(--ink-muted)] mb-1">类目</label>
              <select
                v-model="newGearCategory"
                class="w-full px-2.5 py-1.5 text-xs font-mono-archive bg-[var(--bg-archive)] border border-[var(--border-subtle)] rounded-xs outline-none"
              >
                <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
              </select>
            </div>

            <div>
              <label class="block text-[11px] font-mono-archive text-[var(--ink-muted)] mb-1">设备名称 *</label>
              <input
                v-model="newGearName"
                type="text"
                placeholder="例如: Contax T2"
                class="w-full px-2.5 py-1.5 text-xs font-mono-archive bg-[var(--bg-archive)] border border-[var(--border-subtle)] rounded-xs outline-none"
              />
            </div>

            <div>
              <label class="block text-[11px] font-mono-archive text-[var(--ink-muted)] mb-1">状态标签</label>
              <input
                v-model="newGearStatus"
                type="text"
                placeholder="例如: 随身主力"
                class="w-full px-2.5 py-1.5 text-xs font-mono-archive bg-[var(--bg-archive)] border border-[var(--border-subtle)] rounded-xs outline-none"
              />
            </div>
          </div>

          <div>
            <label class="block text-[11px] font-mono-archive text-[var(--ink-muted)] mb-1">具体规格 / 型号说明</label>
            <input
              v-model="newGearSpec"
              type="text"
              placeholder="例如: 38mm f/2.8 Carl Zeiss T*"
              class="w-full px-2.5 py-1.5 text-xs font-mono-archive bg-[var(--bg-archive)] border border-[var(--border-subtle)] rounded-xs outline-none"
            />
          </div>

          <div>
            <label class="block text-[11px] font-mono-archive text-[var(--ink-muted)] mb-1">使用感受 / 体验笔记</label>
            <input
              v-model="newGearNote"
              type="text"
              placeholder="例如: 快门清脆，成色完好"
              class="w-full px-2.5 py-1.5 text-xs font-mono-archive bg-[var(--bg-archive)] border border-[var(--border-subtle)] rounded-xs outline-none"
            />
          </div>

          <div class="flex justify-end pt-2">
            <button
              @click="handleAddGear"
              type="button"
              class="px-4 py-1.5 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] rounded-xs cursor-pointer hover:opacity-90"
            >
              确认添入列表
            </button>
          </div>
        </div>

        <!-- 设备列表渲染 -->
        <div class="divide-y divide-[var(--border-subtle)]/70 max-h-[46vh] overflow-y-auto pr-1">
          <div
            v-for="(item, idx) in form.gearList"
            :key="item.id || idx"
            class="py-3 flex items-start justify-between gap-4 group"
          >
            <div class="space-y-1">
              <div class="flex items-center gap-2">
                <span class="text-[10px] font-mono-archive px-1.5 py-0.2 rounded-xs border border-[var(--border-subtle)] text-[var(--ink-muted)]">
                  {{ item.category }}
                </span>
                <span class="font-serif-editorial text-base text-[var(--ink-primary)] font-medium">
                  {{ item.name }}
                </span>
                <span v-if="item.status" class="text-[11px] font-mono-archive text-[var(--accent-warm)]">
                  · {{ item.status }}
                </span>
              </div>
              <p v-if="item.spec" class="text-xs font-mono-archive text-[var(--ink-muted)]">
                {{ item.spec }}
              </p>
              <p v-if="item.note" class="text-xs font-serif-editorial text-[var(--ink-secondary)] italic">
                “{{ item.note }}”
              </p>
            </div>

            <button
              @click="removeGear(idx)"
              class="text-[var(--ink-muted)] hover:text-red-700 p-1 cursor-pointer opacity-70 group-hover:opacity-100 transition-opacity"
              title="删除此项"
            >
              <Trash2 class="w-3.5 h-3.5" />
            </button>
          </div>
        </div>
      </div>

      <!-- 底部控制栏 -->
      <div class="pt-4 border-t border-[var(--border-subtle)] flex items-center justify-between">
        <button
          @click="handleResetDefault"
          type="button"
          class="inline-flex items-center gap-1 text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
        >
          <RotateCcw class="w-3.5 h-3.5" />
          <span>恢复预置数据</span>
        </button>

        <div class="flex items-center gap-3">
          <button
            @click="emit('close')"
            type="button"
            class="px-4 py-2 text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
          >
            取消
          </button>
          <button
            @click="handleSave"
            type="button"
            class="px-5 py-2 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] hover:opacity-90 rounded-xs transition-opacity cursor-pointer flex items-center gap-1.5"
          >
            <Check class="w-3.5 h-3.5" />
            <span>保存资料</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
