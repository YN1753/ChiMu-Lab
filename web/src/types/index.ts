export interface Profile {
  id: number
  name: string
  title: string
  bio: string
  avatar: string
  github: string
  email: string
  location: string
  skills: string
}

export interface Project {
  id: number
  title: string
  subtitle: string
  description: string
  category: 'core' | 'lab' | 'tool' | 'study' | string
  tags: string
  demo_url: string
  github_url: string
  status: 'Active' | 'Stable' | 'WIP' | 'Archived' | string
  featured: boolean
  order: number
}

export interface Activity {
  id: number
  date: string
  type: 'commit' | 'release' | 'study' | 'milestone' | string
  title: string
  description: string
  repo_name: string
  count: number
  link: string
}

export interface Stats {
  total_projects: number
  total_activities: number
  active_days: number
  uptime_hours: number
  last_updated: string
}

export interface SiteConfig {
  id: number
  site_name: string
  site_desc: string
  domain: string
  icp_number: string
  icp_link: string
  police_number: string
  police_code: string
}
