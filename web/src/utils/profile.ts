import type { UserProfile } from '../types'

const PROFILE_STORAGE_KEY = 'chimu_curator_profile'

export const DEFAULT_PROFILE: UserProfile = {
  name: '迟暮',
  title: '独立开发者 · 胶片摄影与手冲咖啡学徒 · 城市漫游者',
  avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=800&q=80',
  location: '中国 · 杭州 (西湖区 / 余杭)',
  philosophy: '生活大于项目，真实大于表演。',
  bio: '我喜欢具体的、有体温的事物。清晨手冲咖啡时漫开的花果香气、骑着公路车穿过西湖杨公堤时迎面扑来的湿润晚风、老旁轴相机快门闭合时清脆的一声咔哒，以及在深夜书房里用 Go 雕琢出一套干净的并发状态机。生活终究是由一顿顿具体的饭、一趟趟具体的骑行、一次次真实的思索拼接而成的。',
  email: 'chimu@codeactivityhub.top',
  github: 'https://github.com/YN1753',
  gearList: [
    // 1. 影像与暗房
    {
      id: 'gear-1',
      category: '影像与暗房',
      name: 'Contax T2',
      spec: 'Titanium Silver · Carl Zeiss Sonnar 38mm f/2.8 T*',
      status: '随身主力机',
      note: '极简旁轴机械手感，锐利且通透的德系发色，记录杭州四季街头。',
    },
    {
      id: 'gear-2',
      category: '影像与暗房',
      name: 'Kodak Portra 400',
      spec: '135 彩色负片 · ISO 400',
      status: '日常主力卷',
      note: '柔和温暖的肤色过渡，阴天与室内抓拍必备底片。',
    },
    {
      id: 'gear-3',
      category: '影像与暗房',
      name: 'Kodak Ektar 100',
      spec: '135 彩色负片 · ISO 100',
      status: '风光特选卷',
      note: '极细银盐微颗粒，晴日高反差与浓郁色彩表现力。',
    },
    {
      id: 'gear-4',
      category: '影像与暗房',
      name: 'Ilford HP5 Plus 400',
      spec: '135 黑白银盐胶片 · ISO 400',
      status: '纯银盐纪实',
      note: '雨天与深夜街头的宽容度之王，颗粒粗粝而坦诚。',
    },

    // 2. 生产力工作台
    {
      id: 'gear-5',
      category: '工作台与算力',
      name: 'MacBook Pro 14"',
      spec: 'Apple Silicon M-series · 32GB RAM · 深空灰',
      status: '核心生产力',
      note: '全天续航与安静无声的风扇，随时随地开启移动书房。',
    },
    {
      id: 'gear-6',
      category: '工作台与算力',
      name: '27" 4K 广色域显示器',
      spec: 'IPS Black 面板 · 98% DCI-P3 · Type-C 90W 反向供电',
      status: '桌面视界',
      note: '高精度修图、色彩校准与分屏多栏写代码。',
    },
    {
      id: 'gear-7',
      category: '工作台与算力',
      name: '客制化机械键盘',
      spec: '静音红轴 · 灰白复古 PBT 键帽 · Gasket 结构',
      status: '码字利器',
      note: '安静沉稳的敲击反馈，深夜敲代码不打扰宁静。',
    },
    {
      id: 'gear-8',
      category: '工作台与算力',
      name: 'MCHOSE A7 无线鼠标',
      spec: 'PAW3395 传感器 · 超轻量化手感',
      status: '日常双模',
      note: '握持轻盈，替换了用了四年的旧设备。',
    },

    // 3. 手冲咖啡工坊
    {
      id: 'gear-9',
      category: '手冲咖啡工坊',
      name: 'Hario V60 01 滤杯',
      spec: '有田烧纯白陶瓷 · 经典螺旋导流肋骨',
      status: '晨间仪式',
      note: '大口径流速可控，浅烘豆花果香气萃取的标准器材。',
    },
    {
      id: 'gear-10',
      category: '手冲咖啡工坊',
      name: '泰摩栗子 C3 手磨',
      spec: 'S2C 660 双五星不锈钢磨芯',
      status: '精密研磨',
      note: '极低细粉率，手摇研磨的阻尼感是清晨醒脑的最佳动作。',
    },
    {
      id: 'gear-11',
      category: '手冲咖啡工坊',
      name: 'Fellow Stagg EKG 手冲壶',
      spec: '900ml · 亚光黑 · 细口鹅颈设计',
      status: '温度掌控',
      note: '92℃ 阶梯水温定点垂直匀速注水，拒绝玄学冲煮。',
    },
    {
      id: 'gear-12',
      category: '手冲咖啡工坊',
      name: '埃塞俄比亚日晒花魁 G1',
      spec: '罕贝拉产区 · 浅度烘焙',
      status: '常备豆单',
      note: '带有明亮柑橘、白桃乌龙与茉莉花香的干净甜感。',
    },

    // 4. 夜骑巡航
    {
      id: 'gear-13',
      category: '夜骑巡航',
      name: '破风公路车',
      spec: '700c 铝合金轻量车架 · 碳纤维前叉 · 禧玛诺套件',
      status: '西湖破风战车',
      note: '杨公堤、茅家埠与龙井路 21KM 深夜独行巡航。',
    },
    {
      id: 'gear-14',
      category: '夜骑巡航',
      name: 'Garmin 佳明 GPS 码表',
      spec: 'ANT+ 双模心率与踏频同步 · 高精度气压计',
      status: '节奏监控',
      note: '不看配速，只看呼吸节奏与踏频自律。',
    },

    // 5. 纸笔与日常携带
    {
      id: 'gear-15',
      category: '纸笔与日常',
      name: 'Midori MD Notebook (A5)',
      spec: '方格本 · 日本 MD 特种草本纸',
      status: '随身手账',
      note: '温润的米白纸面，钢笔书写不洇不透，记录电光石火的想法。',
    },
    {
      id: 'gear-16',
      category: '纸笔与日常',
      name: 'Pilot 78G 钢笔',
      spec: 'F 细尖 · 搭配鲶鱼永恒黑墨水',
      status: '日常书写',
      note: '出水均匀节制，落笔即成确凿记忆。',
    },
  ],
}

export function loadCuratorProfile(): UserProfile {
  try {
    const raw = localStorage.getItem(PROFILE_STORAGE_KEY)
    if (raw) {
      const parsed = JSON.parse(raw)
      return {
        ...DEFAULT_PROFILE,
        ...parsed,
        gearList: Array.isArray(parsed.gearList) ? parsed.gearList : DEFAULT_PROFILE.gearList,
      }
    }
  } catch {
    // ignore
  }
  return { ...DEFAULT_PROFILE }
}

export function saveCuratorProfile(profile: UserProfile): void {
  try {
    localStorage.setItem(PROFILE_STORAGE_KEY, JSON.stringify(profile))
    window.dispatchEvent(new CustomEvent('chimu-profile-updated', { detail: profile }))
  } catch (err) {
    console.error('Failed to save curator profile to localStorage:', err)
  }
}
