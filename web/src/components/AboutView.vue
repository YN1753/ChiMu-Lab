<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ArrowUpRight, Edit3, MapPin, Mail, Github, Maximize2 } from 'lucide-vue-next'
import type { UserProfile } from '../types'
import { loadCuratorProfile } from '../utils/profile'
import ProfileEditModal from './ProfileEditModal.vue'
import EditorialLightbox from './EditorialLightbox.vue'

const profile = ref<UserProfile>(loadCuratorProfile())
const isEditOpen = ref(false)
const isLightboxOpen = ref(false)
const selectedCategory = ref<string>('all')

const categories = computed(() => {
  const set = new Set<string>()
  profile.value.gearList.forEach((g) => {
    if (g.category) set.add(g.category)
  })
  return Array.from(set)
})

const filteredGear = computed(() => {
  if (selectedCategory.value === 'all') return profile.value.gearList
  return profile.value.gearList.filter((g) => g.category === selectedCategory.value)
})

const handleProfileUpdated = (updated: UserProfile) => {
  profile.value = updated
}

onMounted(() => {
  profile.value = loadCuratorProfile()
  window.addEventListener('chimu-profile-updated', ((e: CustomEvent<UserProfile>) => {
    if (e.detail) profile.value = e.detail
  }) as EventListener)
})
</script>

<template>
  <div class="max-w-4xl mx-auto px-5 sm:px-8 py-20 sm:py-28 text-left space-y-16">
    
    <!-- 标头与编辑入口 -->
    <header class="pb-10 border-b border-[var(--border-subtle)] space-y-3">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
            主理人
          </h1>
          <span class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
            // CURATOR & GEAR
          </span>
        </div>

        <button
          @click="isEditOpen = true"
          class="inline-flex items-center gap-1.5 px-3 py-1.5 text-xs font-mono-archive border border-[var(--border-subtle)] hover:border-[var(--ink-primary)] rounded-xs transition-colors cursor-pointer text-[var(--ink-secondary)] hover:text-[var(--ink-primary)]"
          title="自定义头像、简介与设备清单"
        >
          <Edit3 class="w-3.5 h-3.5" />
          <span>编辑资料与装备</span>
        </button>
      </div>

      <p class="font-serif-editorial text-sm sm:text-base text-[var(--ink-secondary)]">
        关于生活在此处的主理人、随身器物与这本散页档案。
      </p>
    </header>

    <!-- 核心板块 1：主理人装裱肖像与个人叙事 -->
    <section class="grid grid-cols-1 md:grid-cols-12 gap-10 items-start">
      
      <!-- 左栏：胶片装裱肖像与极简身份信息条 -->
      <div class="md:col-span-5 space-y-6">
        <!-- 暗房装裱底片框肖像 -->
        <div
          class="film-frame bg-[var(--bg-subtle)]/80 p-3 rounded-xs border border-[var(--border-subtle)] cursor-zoom-in group"
          @click="isLightboxOpen = true"
          title="点击在暗房灯箱中查看肖像"
        >
          <div class="relative overflow-hidden rounded-xs bg-[var(--bg-archive)] aspect-square">
            <img
              :src="profile.avatar"
              :alt="profile.name"
              class="w-full h-full object-cover group-hover:scale-[1.02] transition-transform duration-500"
              onerror="this.src='https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=800&q=80'"
            />
            <div class="absolute bottom-2 right-2 px-2 py-1 bg-black/60 backdrop-blur-xs text-white text-[10px] font-mono-archive rounded-xs opacity-0 group-hover:opacity-100 transition-opacity flex items-center gap-1">
              <Maximize2 class="w-3 h-3" />
              <span>灯箱大图</span>
            </div>
          </div>

          <!-- 胶片装裱条底款 -->
          <div class="pt-2.5 px-1 flex items-center justify-between text-[11px] font-mono-archive text-[var(--ink-muted)]">
            <span class="tracking-widest uppercase">EXP. 01 // PORTRAIT</span>
            <span>{{ profile.location.split('·')[1]?.trim() || 'HANGZHOU' }}</span>
          </div>
        </div>

        <!-- 身份速览卡片 -->
        <div class="p-4 rounded-xs border border-[var(--border-subtle)] bg-[var(--bg-subtle)]/30 space-y-3 font-mono-archive text-xs">
          <div>
            <div class="text-[10px] uppercase tracking-wider text-[var(--ink-muted)]">// 名字 / CURATOR</div>
            <div class="font-serif-editorial text-lg text-[var(--ink-primary)] font-medium pt-0.5">
              {{ profile.name }}
            </div>
            <div class="text-[11px] text-[var(--ink-muted)] pt-0.5">
              {{ profile.title }}
            </div>
          </div>

          <div class="pt-2 border-t border-[var(--border-subtle)]/60 space-y-1.5 text-[var(--ink-secondary)]">
            <div class="flex items-center gap-1.5">
              <MapPin class="w-3.5 h-3.5 text-[var(--accent-warm)] shrink-0" />
              <span>{{ profile.location }}</span>
            </div>
            <div class="flex items-center gap-1.5">
              <Mail class="w-3.5 h-3.5 text-[var(--accent-warm)] shrink-0" />
              <a :href="`mailto:${profile.email}`" class="hover-underline">{{ profile.email }}</a>
            </div>
            <div class="flex items-center gap-1.5">
              <Github class="w-3.5 h-3.5 text-[var(--accent-warm)] shrink-0" />
              <a :href="profile.github" target="_blank" rel="noopener noreferrer" class="hover-underline inline-flex items-center gap-1">
                <span>GitHub @YN1753</span>
                <ArrowUpRight class="w-3 h-3" />
              </a>
            </div>
          </div>
        </div>
      </div>

      <!-- 右栏：个人独白与建站哲学 -->
      <div class="md:col-span-7 space-y-8 font-serif-editorial text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed">
        
        <!-- 核心信条引言 -->
        <blockquote class="text-xl sm:text-2xl text-[var(--ink-primary)] italic border-l-2 border-[var(--accent-warm)] pl-5 py-1">
          “{{ profile.philosophy }}”
        </blockquote>

        <!-- 关于我 -->
        <section class="space-y-3">
          <h3 class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
            // 自白 · MONOLOGUE
          </h3>
          <p class="whitespace-pre-line text-[var(--ink-primary)]">
            {{ profile.bio }}
          </p>
        </section>

        <!-- 为什么做这本《散页》 -->
        <section class="space-y-3 pt-4 border-t border-[var(--border-subtle)]">
          <h3 class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
            // 为什么做《散页》 · WHY LOOSE LEAF
          </h3>
          <p>
            在当下的互联网上，很多人的个人主页变成了千篇一律的职业履历展厅：满屏的技能徽章、KPI 指标、格式化的项目卡片与给招聘者看的说辞。
          </p>
          <p>
            但我渐渐发现，那些真正构成我生命质感的东西——读完一本书时的触动、和朋友在深夜连麦打游戏的畅快、秋天满觉陇落满青石板的金桂底片——在那些标准的简历模板里根本无处安放。
          </p>
          <p>
            所以我决定推翻传统模板，把这里做成一本<strong>散页生活档案</strong>。时间是这里唯一的刻度。无论是一张底片、一段随笔、一个亲手造的轮子，还是一包咖啡豆的萃取数据，它们都是我生命中真真切切发生过的印记。
          </p>
        </section>

      </div>

    </section>

    <!-- 核心板块 2：随身与桌面设备清单 (Desk, Gear & EDC) -->
    <section class="space-y-8 pt-8 border-t border-[var(--border-subtle)]">
      
      <div class="flex flex-col sm:flex-row sm:items-baseline justify-between gap-4">
        <div>
          <div class="flex items-center gap-2">
            <span class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
              // 随身与桌面器物
            </span>
            <span class="font-mono-archive text-xs uppercase tracking-widest text-[var(--accent-warm)]">
              · DESK & CARRY (EDC)
            </span>
          </div>
          <h2 class="font-serif-editorial text-2xl sm:text-3xl font-normal text-[var(--ink-primary)] mt-1">
            陪伴日常的手艺工具与装备
          </h2>
        </div>

        <span class="text-xs font-mono-archive text-[var(--ink-muted)]">
          共收录 {{ profile.gearList.length }} 项设备器物
        </span>
      </div>

      <!-- 设备类别筛选切换栏 -->
      <div
        v-if="categories.length > 1"
        class="flex flex-wrap items-center gap-x-6 gap-y-2 text-xs font-mono-archive text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-3"
      >
        <span class="text-[var(--ink-secondary)] mr-1">分类 //</span>
        <button
          @click="selectedCategory = 'all'"
          class="transition-colors cursor-pointer py-1 hover:text-[var(--ink-primary)]"
          :class="{
            'text-[var(--ink-primary)] font-semibold border-b border-[var(--ink-primary)]':
              selectedCategory === 'all',
          }"
        >
          全部器物 ({{ profile.gearList.length }})
        </button>
        <button
          v-for="cat in categories"
          :key="cat"
          @click="selectedCategory = cat"
          class="transition-colors cursor-pointer py-1 hover:text-[var(--ink-primary)]"
          :class="{
            'text-[var(--ink-primary)] font-semibold border-b border-[var(--ink-primary)]':
              selectedCategory === cat,
          }"
        >
          {{ cat }}
        </button>
      </div>

      <!-- 设备器物清单卡片 (Editorial 纸本图鉴排版) -->
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
        <article
          v-for="gear in filteredGear"
          :key="gear.id || gear.name"
          class="p-5 rounded-xs border border-[var(--border-subtle)] bg-[var(--bg-subtle)]/40 hover:border-[var(--border-divider)] transition-all space-y-3"
        >
          <!-- 类目标签与状态 -->
          <div class="flex items-center justify-between text-[11px] font-mono-archive">
            <span class="text-[var(--ink-muted)] uppercase tracking-wider">
              {{ gear.category }}
            </span>
            <span
              v-if="gear.status"
              class="border border-[var(--border-subtle)] px-2 py-0.5 rounded-xs text-[10px] text-[var(--accent-warm)]"
            >
              {{ gear.status }}
            </span>
          </div>

          <!-- 设备名称与规格 -->
          <div class="space-y-1">
            <h3 class="font-serif-editorial text-xl font-medium text-[var(--ink-primary)] tracking-tight">
              {{ gear.name }}
            </h3>
            <p v-if="gear.spec" class="font-mono-archive text-xs text-[var(--ink-muted)]">
              {{ gear.spec }}
            </p>
          </div>

          <!-- 使用感受 / 备注 -->
          <p v-if="gear.note" class="font-serif-editorial text-sm text-[var(--ink-secondary)] leading-relaxed italic border-l border-[var(--border-subtle)] pl-3 py-0.5">
            “{{ gear.note }}”
          </p>
        </article>
      </div>

    </section>

    <!-- 底部联络区 -->
    <footer class="pt-10 border-t border-[var(--border-subtle)] font-mono-archive text-xs text-[var(--ink-muted)] flex flex-wrap items-center justify-between gap-4">
      <span>来信交流：{{ profile.email }}</span>
      <a
        :href="profile.github"
        target="_blank"
        rel="noopener noreferrer"
        class="text-[var(--ink-primary)] hover-underline inline-flex items-center gap-1"
      >
        <span>GitHub @YN1753</span>
        <ArrowUpRight class="w-3 h-3" />
      </a>
    </footer>

    <!-- 编辑资料弹窗 -->
    <ProfileEditModal
      :isOpen="isEditOpen"
      :profile="profile"
      @close="isEditOpen = false"
      @updated="handleProfileUpdated"
    />

    <!-- 暗房肖像大图灯箱 -->
    <EditorialLightbox
      :isOpen="isLightboxOpen"
      :imageUrl="profile.avatar"
      :title="`${profile.name} · 主理人肖像`"
      :meta="'35mm PORTRAIT ARCHIVE'"
      :location="profile.location"
      :date="'2026 档案收录'"
      @close="isLightboxOpen = false"
    />

  </div>
</template>
