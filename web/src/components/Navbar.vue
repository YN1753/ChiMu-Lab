<script setup lang="ts">
import { ref } from 'vue'
import { Terminal, Github, ExternalLink, Menu, X } from 'lucide-vue-next'

const mobileMenuOpen = ref(false)

const navLinks = [
  { name: '首页', href: '#hero' },
  { name: '作品展厅', href: '#projects' },
  { name: '代码动态', href: '#activity' },
  { name: '关于', href: '#about' },
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
  <header class="sticky top-0 z-50 w-full border-b border-white/10 bg-slate-950/80 backdrop-blur-md">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 h-16 flex items-center justify-between">
      <!-- Logo -->
      <a href="#hero" class="flex items-center gap-3 group">
        <div class="w-10 h-10 rounded-xl bg-gradient-to-tr from-cyan-500 to-blue-600 flex items-center justify-center text-white font-bold text-lg shadow-lg shadow-cyan-500/20 group-hover:scale-105 transition-transform">
          <Terminal class="w-5 h-5 text-white" />
        </div>
        <div>
          <span class="font-bold text-lg tracking-tight text-white flex items-center gap-1.5">
            ChiMu-Lab
            <span class="text-xs px-2 py-0.5 rounded-full bg-cyan-950 text-cyan-400 border border-cyan-800 font-mono">v1.0</span>
          </span>
          <p class="text-xs text-slate-400 font-mono -mt-1 hidden sm:block">迟暮实验室 · Code Activity Hub</p>
        </div>
      </a>

      <!-- Desktop Nav -->
      <nav class="hidden md:flex items-center gap-8">
        <button
          v-for="item in navLinks"
          :key="item.name"
          @click="scrollToSection(item.href)"
          class="text-sm font-medium text-slate-300 hover:text-cyan-400 transition-colors cursor-pointer"
        >
          {{ item.name }}
        </button>
      </nav>

      <!-- Right Actions -->
      <div class="hidden md:flex items-center gap-4">
        <div class="flex items-center gap-2 px-3 py-1 rounded-full bg-slate-900 border border-slate-800 text-xs font-mono text-emerald-400">
          <span class="w-2 h-2 rounded-full bg-emerald-400 animate-pulse"></span>
          Server Operational
        </div>

        <a
          href="https://github.com/YN1753/ChiMu-Lab"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-2 px-3.5 py-1.5 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-200 text-sm font-medium border border-slate-700 hover:border-slate-600 transition-all shadow-sm"
        >
          <Github class="w-4 h-4" />
          <span>GitHub</span>
          <ExternalLink class="w-3.5 h-3.5 text-slate-400" />
        </a>
      </div>

      <!-- Mobile Menu Button -->
      <button
        @click="mobileMenuOpen = !mobileMenuOpen"
        class="md:hidden p-2 rounded-lg text-slate-400 hover:text-white hover:bg-slate-900"
      >
        <Menu v-if="!mobileMenuOpen" class="w-6 h-6" />
        <X v-else class="w-6 h-6" />
      </button>
    </div>

    <!-- Mobile Nav Dropdown -->
    <div v-if="mobileMenuOpen" class="md:hidden border-b border-white/10 bg-slate-950/95 px-4 py-4 space-y-3">
      <button
        v-for="item in navLinks"
        :key="item.name"
        @click="scrollToSection(item.href)"
        class="block w-full text-left py-2 text-base font-medium text-slate-300 hover:text-cyan-400"
      >
        {{ item.name }}
      </button>
      <div class="pt-3 border-t border-slate-800 flex items-center justify-between">
        <span class="text-xs text-emerald-400 font-mono flex items-center gap-2">
          <span class="w-2 h-2 rounded-full bg-emerald-400"></span>
          系统正常运行
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
