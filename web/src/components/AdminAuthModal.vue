<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { X, Key, Check, CheckCircle2 } from 'lucide-vue-next'
import { getAdminToken, setAdminToken, removeAdminToken } from '../utils/api'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const currentKey = ref('')
const inputKey = ref('')
const savedSuccess = ref(false)

const loadKey = () => {
  currentKey.value = getAdminToken()
  inputKey.value = currentKey.value
  savedSuccess.value = false
}

onMounted(() => {
  loadKey()
})

const handleSave = () => {
  if (inputKey.value.trim()) {
    setAdminToken(inputKey.value.trim())
    currentKey.value = inputKey.value.trim()
  } else {
    removeAdminToken()
    currentKey.value = ''
  }
  savedSuccess.value = true
  setTimeout(() => {
    savedSuccess.value = false
  }, 2000)
}

const handleClear = () => {
  removeAdminToken()
  currentKey.value = ''
  inputKey.value = ''
  savedSuccess.value = true
  setTimeout(() => {
    savedSuccess.value = false
  }, 2000)
}
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-xs animate-in fade-in duration-200"
    @click.self="emit('close')"
  >
    <div
      class="bg-[var(--bg-archive)] border border-[var(--border-subtle)] shadow-[0_16px_40px_rgba(0,0,0,0.1)] w-full max-w-md rounded-xs p-6 sm:p-8 text-left space-y-6"
    >
      <div class="flex items-center justify-between pb-4 border-b border-[var(--border-subtle)]">
        <div class="flex items-center gap-2.5">
          <Key class="w-4 h-4 text-[var(--ink-secondary)]" />
          <h2 class="font-serif-editorial text-xl font-medium text-[var(--ink-primary)]">
            暗房钥匙 // ADMIN KEY
          </h2>
        </div>
        <button
          @click="emit('close')"
          class="p-1 text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
        >
          <X class="w-4 h-4" />
        </button>
      </div>

      <div class="space-y-4 font-serif-editorial">
        <p class="text-xs text-[var(--ink-secondary)] leading-relaxed">
          若服务器配置了 <code>ADMIN_API_KEY</code>，在此输入你的管理员密钥。输入后将保存在本地浏览器中，后续新增生活记录、记账与上传时将自动携带鉴权。
        </p>

        <div class="space-y-1.5">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">管理员密钥 (Token)</label>
          <input
            v-model="inputKey"
            type="password"
            placeholder="输入在 .env 中设置的 ADMIN_API_KEY"
            class="w-full px-3 py-2 text-sm font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
          />
        </div>

        <div class="flex items-center justify-between pt-1 text-xs font-mono-archive">
          <span class="text-[var(--ink-muted)]">当前状态：</span>
          <span v-if="currentKey" class="text-emerald-700 flex items-center gap-1">
            <CheckCircle2 class="w-3.5 h-3.5" />
            <span>已保存本地钥匙</span>
          </span>
          <span v-else class="text-[var(--ink-muted)]">
            未配置本地钥匙
          </span>
        </div>

        <div v-if="savedSuccess" class="p-2.5 bg-emerald-50 text-emerald-800 text-xs rounded-xs border border-emerald-200 flex items-center gap-1.5 font-mono-archive">
          <Check class="w-3.5 h-3.5" />
          <span>钥匙状态已更新！</span>
        </div>
      </div>

      <div class="pt-2 flex items-center justify-between border-t border-[var(--border-subtle)]">
        <button
          v-if="currentKey"
          @click="handleClear"
          type="button"
          class="text-xs font-mono-archive text-red-600/80 hover:text-red-700 cursor-pointer"
        >
          清除钥匙
        </button>
        <span v-else></span>

        <div class="flex items-center gap-2">
          <button
            @click="emit('close')"
            class="px-3.5 py-1.5 text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
          >
            完成
          </button>
          <button
            @click="handleSave"
            class="px-4 py-1.5 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] rounded-xs cursor-pointer hover:opacity-90 flex items-center gap-1"
          >
            <Check class="w-3.5 h-3.5" />
            <span>保存</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
