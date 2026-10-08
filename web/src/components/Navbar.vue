<script setup lang="ts">
import { ref } from 'vue'
import { Github, Menu, X, ArrowUpRight } from 'lucide-vue-next'

const mobileMenuOpen = ref(false)

const navLinks = [
  { name: '序幕 / Overview', href: '#hero' },
  { name: '项目 / Projects', href: '#projects' },
  { name: '动态 / Activity', href: '#activity' },
]

const scrollToSection = (href: string) => {
  mobileMenuOpen.value = false
  const target = document.querySelector(href)
  if (target) {
    target.scrollIntoView({ behavior: 'smooth' })
  }
}
</script>

<template>
  <header class="sticky top-0 z-50 w-full border-b border-[#e8e6df] bg-[#f7f6f2]/85 backdrop-blur-xl">
    <div class="max-w-6xl mx-auto px-5 sm:px-8 h-18 flex items-center justify-between">
      <!-- 品牌 Logo -->
      <a href="#hero" class="flex items-center gap-3.5 group">
        <div class="w-8 h-8 rounded-lg bg-[#14151a] flex items-center justify-center text-white font-serif-cinematic text-sm font-semibold shadow-sm group-hover:scale-105 transition-transform">
          迟
        </div>
        <div class="text-left">
          <div class="flex items-center gap-2">
            <span class="font-bold text-sm tracking-tight text-[#14151a] font-serif-cinematic text-base">
              ChiMu-Lab
            </span>
            <span class="text-[10px] px-1.5 py-0.2 rounded-full bg-[#edeae1] text-[#716e64] font-mono">
              Personal Workbench
            </span>
          </div>
          <p class="text-[11px] text-[#8c8f9b] font-mono tracking-wider -mt-0.5 hidden sm:block">
            迟暮的自留地 · codeactivityhub.top
          </p>
        </div>
      </a>

      <!-- 桌面端导航 -->
      <nav class="hidden md:flex items-center gap-1 bg-[#edeae1]/60 border border-[#e2ded4] rounded-full px-4 py-1.5 backdrop-blur-sm">
        <button
          v-for="item in navLinks"
          :key="item.name"
          @click="scrollToSection(item.href)"
          class="text-xs font-medium text-[#525662] hover:text-[#14151a] px-3.5 py-1 rounded-full hover:bg-white/80 transition-all cursor-pointer font-sans"
        >
          {{ item.name }}
        </button>
      </nav>

      <!-- 右侧：真实状态与 GitHub -->
      <div class="hidden md:flex items-center gap-3.5">
        <div class="flex items-center gap-2 px-3 py-1.5 rounded-full bg-white border border-[#e8e6df] text-xs font-mono text-[#0d766e] shadow-2xs">
          <span class="w-1.5 h-1.5 rounded-full bg-[#0d766e] animate-pulse"></span>
          <span class="text-[11px] tracking-wide">CVM Port 443 Live</span>
        </div>

        <a
          href="https://github.com/YN1753"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-[#14151a] hover:bg-[#272932] text-white text-xs font-medium transition-all shadow-sm group"
        >
          <Github class="w-3.5 h-3.5" />
          <span>YN1753</span>
          <ArrowUpRight class="w-3 h-3 text-slate-400 group-hover:text-white transition-colors" />
        </a>
      </div>

      <!-- 移动端汉堡按钮 -->
      <button
        @click="mobileMenuOpen = !mobileMenuOpen"
        class="md:hidden p-2 rounded-lg text-[#525662] hover:text-[#14151a] hover:bg-black/5"
      >
        <Menu v-if="!mobileMenuOpen" class="w-5 h-5" />
        <X v-else class="w-5 h-5" />
      </button>
    </div>

    <!-- 移动端下拉 -->
    <div v-if="mobileMenuOpen" class="md:hidden border-b border-[#e8e6df] bg-[#f7f6f2] px-6 py-5 space-y-3">
      <button
        v-for="item in navLinks"
        :key="item.name"
        @click="scrollToSection(item.href)"
        class="block w-full text-left py-2 text-sm font-medium text-[#525662] hover:text-[#14151a]"
      >
        {{ item.name }}
      </button>
      <div class="pt-3 border-t border-[#e8e6df] flex items-center justify-between">
        <span class="text-xs text-[#0d766e] font-mono flex items-center gap-2">
          <span class="w-1.5 h-1.5 rounded-full bg-[#0d766e]"></span>
          服务稳定在线
        </span>
        <a
          href="https://github.com/YN1753"
          target="_blank"
          class="text-xs text-[#14151a] font-semibold flex items-center gap-1 font-mono"
        >
          GitHub @YN1753 <ArrowUpRight class="w-3 h-3" />
        </a>
      </div>
    </div>
  </header>
</template>
