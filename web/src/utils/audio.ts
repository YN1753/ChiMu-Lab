// Web Audio API 动态合成微音效与自然白噪音，零外部资源依赖

class AudioEngine {
  private ctx: AudioContext | null = null
  private rainNode: AudioNode | null = null
  private rainGain: GainNode | null = null
  public soundEnabled: boolean = true
  public ambientRainEnabled: boolean = false

  private initContext() {
    if (!this.ctx) {
      const AudioCtx = window.AudioContext || (window as unknown as { webkitAudioContext: typeof AudioContext }).webkitAudioContext
      if (AudioCtx) {
        this.ctx = new AudioCtx()
      }
    }
    if (this.ctx && this.ctx.state === 'suspended') {
      this.ctx.resume()
    }
  }

  // 机械相机快门声 (Mechanical Shutter Click)
  public playShutter() {
    if (!this.soundEnabled) return
    try {
      this.initContext()
      if (!this.ctx) return

      const now = this.ctx.currentTime

      // 第一次快门释放咔哒声
      this.createClick(now, 1200, 0.015, 0.15)
      // 第二次帘幕闭合咔哒声
      this.createClick(now + 0.035, 800, 0.018, 0.12)
    } catch {
      // 浏览器安全策略容错
    }
  }

  // 纸张翻转 / 卡片轻微扑通声 (Paper Flip)
  public playPaperFlip() {
    if (!this.soundEnabled) return
    try {
      this.initContext()
      if (!this.ctx) return

      const now = this.ctx.currentTime
      const osc = this.ctx.createOscillator()
      const gain = this.ctx.createGain()
      const filter = this.ctx.createBiquadFilter()

      osc.type = 'sine'
      osc.frequency.setValueAtTime(140, now)
      osc.frequency.exponentialRampToValueAtTime(45, now + 0.06)

      filter.type = 'lowpass'
      filter.frequency.setValueAtTime(300, now)

      gain.gain.setValueAtTime(0.08, now)
      gain.gain.exponentialRampToValueAtTime(0.001, now + 0.06)

      osc.connect(filter)
      filter.connect(gain)
      gain.connect(this.ctx.destination)

      osc.start(now)
      osc.stop(now + 0.07)
    } catch {
      // 忽略音频异常
    }
  }

  // 轻微金属敲击 / 按钮微触感 (Tink)
  public playTink() {
    if (!this.soundEnabled) return
    try {
      this.initContext()
      if (!this.ctx) return

      const now = this.ctx.currentTime
      const osc = this.ctx.createOscillator()
      const gain = this.ctx.createGain()

      osc.type = 'triangle'
      osc.frequency.setValueAtTime(2400, now)
      osc.frequency.exponentialRampToValueAtTime(1200, now + 0.03)

      gain.gain.setValueAtTime(0.04, now)
      gain.gain.exponentialRampToValueAtTime(0.001, now + 0.03)

      osc.connect(gain)
      gain.connect(this.ctx.destination)

      osc.start(now)
      osc.stop(now + 0.04)
    } catch {
      // 忽略
    }
  }

  // 切换秋夜细雨自然白噪音 (Ambient Rain Synthesizer)
  public toggleRain(): boolean {
    try {
      this.initContext()
      if (!this.ctx) return false

      if (this.ambientRainEnabled) {
        if (this.rainGain && this.ctx) {
          const now = this.ctx.currentTime
          this.rainGain.gain.linearRampToValueAtTime(0.0001, now + 0.5)
          setTimeout(() => {
            if (this.rainNode) {
              this.rainNode.disconnect()
              this.rainNode = null
            }
          }, 600)
        }
        this.ambientRainEnabled = false
        return false
      } else {
        // 生成 3 秒粉红噪声缓冲用于平滑循环
        const bufferSize = this.ctx.sampleRate * 3
        const noiseBuffer = this.ctx.createBuffer(1, bufferSize, this.ctx.sampleRate)
        const output = noiseBuffer.getChannelData(0)
        let b0 = 0, b1 = 0, b2 = 0, b3 = 0, b4 = 0, b5 = 0, b6 = 0
        for (let i = 0; i < bufferSize; i++) {
          const white = Math.random() * 2 - 1
          b0 = 0.99886 * b0 + white * 0.0555179
          b1 = 0.99332 * b1 + white * 0.0750759
          b2 = 0.96900 * b2 + white * 0.1538520
          b3 = 0.86650 * b3 + white * 0.3104856
          b4 = 0.55000 * b4 + white * 0.5329522
          b5 = -0.7616 * b5 - white * 0.0168980
          output[i] = (b0 + b1 + b2 + b3 + b4 + b5 + b6 + white * 0.5362) * 0.035
          b6 = white * 0.115926
        }

        const whiteNoise = this.ctx.createBufferSource()
        whiteNoise.buffer = noiseBuffer
        whiteNoise.loop = true

        const filter = this.ctx.createBiquadFilter()
        filter.type = 'lowpass'
        filter.frequency.setValueAtTime(850, this.ctx.currentTime)

        this.rainGain = this.ctx.createGain()
        this.rainGain.gain.setValueAtTime(0.001, this.ctx.currentTime)
        this.rainGain.gain.linearRampToValueAtTime(0.05, this.ctx.currentTime + 1.2)

        whiteNoise.connect(filter)
        filter.connect(this.rainGain)
        this.rainGain.connect(this.ctx.destination)

        whiteNoise.start()
        this.rainNode = whiteNoise
        this.ambientRainEnabled = true
        return true
      }
    } catch {
      return false
    }
  }

  private createClick(time: number, freq: number, duration: number, volume: number) {
    if (!this.ctx) return
    const osc = this.ctx.createOscillator()
    const gain = this.ctx.createGain()
    const filter = this.ctx.createBiquadFilter()

    osc.type = 'sine'
    osc.frequency.setValueAtTime(freq, time)
    osc.frequency.exponentialRampToValueAtTime(freq * 0.3, time + duration)

    filter.type = 'bandpass'
    filter.frequency.setValueAtTime(freq, time)
    filter.Q.setValueAtTime(3, time)

    gain.gain.setValueAtTime(volume, time)
    gain.gain.exponentialRampToValueAtTime(0.001, time + duration)

    osc.connect(filter)
    filter.connect(gain)
    gain.connect(this.ctx.destination)

    osc.start(time)
    osc.stop(time + duration)
  }
}

export const audio = new AudioEngine()
