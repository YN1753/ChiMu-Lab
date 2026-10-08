<script setup lang="ts">
import { ref } from 'vue'
import { Terminal, Github, ExternalLink, Menu, X } from 'lucide-vue-next'

const mobileMenuOpen = ref(false)

const navLinks = [
  { name: '首页', href: '#hero' },
  { name: '实验室展厅', href: '#projects' },
  { name: '代码动态', href: '#activity' },
  { name: '工程笔记', href: '#notes' },
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
  <header class="sticky top-0 z-50 w-full border-b border-white/[0.06] bg-[#07080c]/85 backdrop-blur-xl">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 h-16 flex items-center justify-between">
      <!-- Brand Logo -->
      <a href="#hero" class="flex items-center gap-3 group">
        <div class="w-9 h-9 rounded-xl bg-gradient-to-tr from-cyan-500 via-sky-500 to-indigo-600 p-[1px] shadow-lg shadow-cyan-500/20 group-hover:shadow-cyan-500/40 transition-all">
          <div class="w-full h-full bg-[#0a0d14] rounded-[11px] flex items-center justify-center">
            <Terminal class="w-4 h-4 text-cyan-400 group-hover:scale-110 transition-transform" />
          </div>
        </div>
        <div class="text-left">
          <div class="flex items-center gap-2">
            <span class="font-bold text-sm sm:text-base tracking-tight text-white group-hover:text-cyan-300 transition-colors">
              ChiMu-Lab
            </span>
            <span class="text-[10px] px-1.5 py-0.5 rounded-full bg-cyan-950/80 text-cyan-400 border border-cyan-800/80 font-mono">
              v2.0
            </span>
          </div>
          <p class="text-[11px] text-slate-500 font-mono hidden sm:block">迟暮的极客工坊 · codeactivityhub.top</p>
        </div>
      </a>

      <!-- Desktop Navigation Links -->
      <nav class="hidden md:flex items-center gap-1 bg-white/[0.03] border border-white/[0.06] rounded-full px-4 py-1.5 backdrop-blur-md">
        <button
          v-for="item in navLinks"
          :key="item.name"
          @click="scrollToSection(item.href)"
          class="text-xs font-medium text-slate-400 hover:text-white px-3 py-1 rounded-full hover:bg-white/[0.06] transition-all cursor-pointer"
        >
          {{ item.name }}
        </button>
      </nav>

      <!-- Right Actions: System Pulse & GitHub -->
      <div class="hidden md:flex items-center gap-3">
        <div class="flex items-center gap-2 px-3 py-1.5 rounded-full bg-[#0d1017] border border-white/[0.08] text-xs font-mono text-emerald-400 shadow-inner">
          <span class="relative flex h-2 w-2">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
            <span class="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
          </span>
          <span class="text-[11px]">System Online</span>
        </div>

        <a
          href="https://github.com/YN1753/ChiMu-Lab"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-white/[0.06] hover:bg-white/[0.1] text-slate-200 text-xs font-medium border border-white/[0.08] hover:border-white/20 transition-all shadow-sm"
        >
          <Github class="w-3.5 h-3.5" />
          <span>GitHub</span>
          <ExternalLink class="w-3 h-3 text-slate-500" />
        </a>
      </div>

      <!-- Mobile Menu Toggle -->
      <button
        @click="mobileMenuOpen = !mobileMenuOpen"
        class="md:hidden p-2 rounded-lg text-slate-400 hover:text-white hover:bg-white/[0.05]"
      >
        <Menu v-if="!mobileMenuOpen" class="w-5 h-5" />
        <X v-else class="w-5 h-5" />
      </button>
    </div>

    <!-- Mobile Navigation Drawer -->
    <div v-if="mobileMenuOpen" class="md:hidden border-b border-white/[0.08] bg-[#07080c]/98 px-5 py-4 space-y-3">
      <button
        v-for="item in navLinks"
        :key="item.name"
        @click="scrollToSection(item.href)"
        class="block w-full text-left py-2 text-sm font-medium text-slate-300 hover:text-cyan-400"
      >
        {{ item.name }}
      </button>
      <div class="pt-3 border-t border-white/[0.08] flex items-center justify-between">
        <span class="text-xs text-emerald-400 font-mono flex items-center gap-2">
          <span class="w-2 h-2 rounded-full bg-emerald-400"></span>
          服务稳定运行
        </span>
        <a
          href="https://github.com/YN1753/ChiMu-Lab"
          target="_blank"
          class="text-xs text-cyan-400 flex items-center gap-1 font-mono"
        >
          GitHub 仓库 <ExternalLink class="w-3 h-3" />
        </a>
      </div>
    </div>
  </header>
</template>
