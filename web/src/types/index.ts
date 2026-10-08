export type EntryType =
  | 'thought'
  | 'photo'
  | 'moment'
  | 'place'
  | 'music'
  | 'book'
  | 'game'
  | 'project'
  | 'purchase'
  | 'coffee'
export type ArchiveCategory = 'all' | 'daily' | 'thought' | 'project' | 'collection'

export const getEntryCategory = (entry: LifeEntry): ArchiveCategory => {
  const t = (entry.type || '').toLowerCase()
  if (t === 'thought' || t === 'idea' || t === 'note' || t === 'essay') {
    return 'thought'
  }
  if (t === 'project' || t === 'code' || t === 'craft') {
    return 'project'
  }
  if (t === 'music' || t === 'book' || t === 'game' || t === 'purchase' || t === 'gear' || t === 'movie' || t === 'device') {
    return 'collection'
  }
  return 'daily'
}

export const getCategoryDisplayName = (cat: ArchiveCategory): string => {
  const map: Record<ArchiveCategory, string> = {
    all: '全部记录',
    daily: '日常',
    thought: '想法',
    project: '项目',
    collection: '收藏',
  }
  return map[cat] || '全部记录'
}

export interface LifeEntry {
  id: number
  date: string
  year: string
  month: string
  day: string
  time: string
  type: EntryType | string
  title: string
  content: string
  images?: string
  location?: string
  meta?: string
  tags?: string
  link?: string
  related_project?: string
  featured?: boolean
}

export interface NowStatus {
  id?: number
  building: string
  learning: string
  playing: string
  listening: string
  reading: string
  thinking: string
  using: string
  location: string
  last_updated: string
}

export interface Project {
  id: number
  title: string
  subtitle: string
  description: string
  story: string
  category: string
  tags: string
  demo_url?: string
  github_url: string
  status: string
  featured: boolean
  order: number
}

export interface YearArchiveStat {
  year: string
  total_count: number
  type_counts: Record<string, number>
  months: string[]
}

export interface SiteConfig {
  site_name: string
  site_desc: string
  domain: string
  icp_number: string
  icp_link: string
}
