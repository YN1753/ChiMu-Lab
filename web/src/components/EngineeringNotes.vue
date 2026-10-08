<script setup lang="ts">
import { ref } from 'vue'
import { BookOpen, Terminal, Code } from 'lucide-vue-next'

const activeTab = ref(0)

const notes = [
  {
    id: 1,
    chapter: 'CHAPTER 01',
    title: '纯 Go SQLite (CGO-Free) 实践：告别交叉编译陷阱',
    tag: '后端架构 / Go 机制',
    date: '2026-10-08',
    summary: '为什么在轻量级全栈服务中，我坚决放弃了传统的 mattn/go-sqlite3，全面拥抱纯 Go 实现的 SQLite？',
    content: `
在 Go 生态中，常用的 SQLite 驱动（如 mattn/go-sqlite3）强依赖 CGO。当在 macOS (darwin/arm64) 本地开发，需要交叉编译到 Linux (linux/amd64) 生产服务器时，CGO 会强行要求配置多架构的 gcc 工具链，且极易因 glibc / musl 版本不匹配引发运行时崩溃。

解决方案：采用基于 modernc.org/sqlite 纯 Go 转译的驱动（结合 glebarez/sqlite 与 GORM）。
优势：
1. 彻底关闭 CGO (CGO_ENABLED=0)，实现毫秒级一键单文件跨平台静态编译；
2. 零外部 libc / so 动态库依赖，直接在 Alpine 或极简 Linux CVM 跑起；
3. 本站实测单库常驻内存仅 ~1.76MB，单次查询时延稳定在 0.2ms 以内。
    `,
    codeSnippet: `// 编译命令：纯 Go 零依赖构建 Linux amd64 二进制
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o chimu-server main.go`,
  },
  {
    id: 2,
    chapter: 'CHAPTER 02',
    title: 'ACM 模式算法推演：相邻约束贪心模型的局部最优证明',
    tag: '算法精炼 / 贪心原型',
    date: '2026-10-06',
    summary: '以 P3817 小A的糖果为例，深入剖析为什么在双盒约束中“优先减少右侧”必然构成全局最优？',
    content: `
题目模型：给定长度为 n 的数组，要求任意相邻两个元素之和不超过 x，求最少减少的数值总和。

贪心策略证明：
从左到右线性扫描，对于当前相邻元素 a[i] 与 a[i+1]：
- 如果 a[i] + a[i+1] > x，我们需要扣减 diff = (a[i] + a[i+1]) - x；
- 若扣减 a[i]：仅对当前 (i-1, i) 和 (i, i+1) 生效，但由于从左往右扫描，(i-1, i) 已经满足，改 a[i] 对未来的约束毫无帮助；
- 若扣减 a[i+1]：不仅修复了当前的超限问题，同时减小了 a[i+1] 在下一轮作为左侧元素时的基数，直接为 (i+1, i+2) 创造了最大的松弛度空间。
结论：优先扣减右盒具备严格的无后效性与贪心选择性质，时间复杂度 O(n)，额外空间 O(1)。
    `,
    codeSnippet: `// 核心贪心决策逻辑 (Go)
for i := 0; i < n-1; i++ {
    sum := a[i] + a[i+1]
    if sum > x {
        diff := sum - x
        cost += diff
        a[i+1] -= diff // 贪心：仅削减右盒
    }
}`,
  },
  {
    id: 3,
    chapter: 'CHAPTER 03',
    title: '生产排障手记：从连接超时到 TLS 1.3 毫秒响应',
    tag: '系统运维 / 网络安全',
    date: '2026-10-01',
    summary: '一次针对域名 403、安全组封锁与 Nginx 301 重定向死循环的完整链路复盘。',
    content: `
现象复盘：
1. 外部访问域名返回连接超时或空白 403，导致全国公安联网备案因“测试空壳页面”被驳回；
2. 根因链条：
   - Nginx 将 80 端口 HTTP 请求 301 强制跳转至 443 HTTPS；
   - 云服务商安全组未放通 443 端口，导致浏览器握手包直接被云网关丢弃；
   - 443 站点的 root 路径误指向了缺少 index.html 的空目录，触发 403 Forbidden。
3. 解决步骤：
   - 在安全组开放 TCP:443 端口；
   - 部署 ChiMu-Lab 生产级守护进程并绑定 Nginx 反代；
   - 注入合规备案标号与工信部超链接，实现公网安全合规上线。
    `,
    codeSnippet: `# 验证 TLS 1.3 握手与 HTTP/2 协议支持
curl -I -k https://codeactivityhub.top
# HTTP/2 200 OK (server: nginx/1.24.0)`,
  },
]
</script>

<template>
  <section id="notes" class="py-24 border-t border-[#e8e6df] bg-[#f7f6f2] relative">
    <div class="max-w-6xl mx-auto px-5 sm:px-8 text-left">
      <!-- 章节导语 (Scene Header) -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-14">
        <div>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[#edeae1] text-[#716e64] text-xs font-mono mb-3">
            <BookOpen class="w-3.5 h-3.5 text-[#d97706]" />
            <span>Act III · Engineering Monograph · 工程札记与架构切片</span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-medium tracking-tight text-[#14151a] font-serif-cinematic">
            技术札记与深度思考
          </h2>
          <p class="text-[#525662] text-sm sm:text-base mt-2 max-w-xl font-normal">
            如同一本装订考究的技术专著，记录真实生产环境的工程取舍、底层陷阱与算法推导。
          </p>
        </div>
      </div>

      <!-- 双栏装订书页排版 -->
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-7">
        <!-- 左侧章节目录 -->
        <div class="lg:col-span-5 space-y-3.5">
          <div
            v-for="(note, idx) in notes"
            :key="note.id"
            @click="activeTab = idx"
            :class="[
              'p-6 rounded-2xl border transition-all cursor-pointer film-card',
              activeTab === idx
                ? 'border-[#d97706]/40 bg-white shadow-md'
                : 'border-[#e8e6df] bg-[#faf9f5] hover:border-[#14151a]/20 hover:bg-white'
            ]"
          >
            <div class="flex items-center justify-between text-[11px] font-mono text-[#8c8f9b] mb-2">
              <span class="text-[#92400e] bg-[#fef3c7] px-2 py-0.5 rounded border border-[#fde68a] font-medium">
                {{ note.tag }}
              </span>
              <span>{{ note.chapter }}</span>
            </div>
            <h4 class="text-base font-bold text-[#14151a] font-serif-cinematic mb-2 leading-snug">
              {{ note.title }}
            </h4>
            <p class="text-xs text-[#525662] line-clamp-2 leading-relaxed">
              {{ note.summary }}
            </p>
          </div>
        </div>

        <!-- 右侧文章正文与代码 -->
        <div class="lg:col-span-7">
          <div class="film-card p-8 sm:p-10 bg-white border border-[#e8e6df] shadow-md h-full flex flex-col justify-between">
            <div>
              <div class="flex items-center justify-between pb-4 border-b border-[#f0ede6] mb-6">
                <span class="text-xs font-mono text-[#92400e] font-medium flex items-center gap-1.5">
                  <Terminal class="w-4 h-4 text-[#d97706]" />
                  {{ notes[activeTab].chapter }} // MONOGRAPH
                </span>
                <span class="text-xs font-mono text-[#8c8f9b]">{{ notes[activeTab].date }}</span>
              </div>

              <h3 class="text-2xl sm:text-3xl font-bold text-[#14151a] font-serif-cinematic mb-5">
                {{ notes[activeTab].title }}
              </h3>

              <div class="text-sm text-[#525662] leading-relaxed space-y-4 whitespace-pre-line font-normal">
                {{ notes[activeTab].content.trim() }}
              </div>
            </div>

            <!-- 代码演示块：温暖雅致的印刷体代码框 -->
            <div class="mt-8 pt-6 border-t border-[#f0ede6]">
              <div class="text-[11px] font-mono text-[#8c8f9b] mb-2 flex items-center gap-1.5">
                <Code class="w-3.5 h-3.5 text-[#d97706]" /> Key Snippet & Command
              </div>
              <pre class="p-4 rounded-xl bg-[#faf9f5] border border-[#e8e6df] font-mono text-xs text-[#14151a] overflow-x-auto leading-relaxed shadow-inner"><code>{{ notes[activeTab].codeSnippet }}</code></pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
