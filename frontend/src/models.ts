export interface Metadata {
  created_at: string;
  updated_at: string;
  source?: string;
}

export interface Official {
  _id: string;
  name: string;
  position: string;
  party?: string;
  bio?: string;
  photo_url?: string;
  metadata: Metadata;
}

export interface StatementSource {
  type: string;
  url?: string;
  title?: string;
  date: string;
  media_outlet?: string;
}

export interface Statement {
  _id: string;
  official_id: string;
  content: string;
  source: StatementSource;
  extracted_by?: string;
  status: string;
  topics?: string[];
  sentiment?: string;
  metadata: {
    created_at: string;
    updated_at: string;
    language?: string;
    ai_confidence?: number;
  };
}

export interface Evidence {
  type: string;
  source: string;
  url?: string;
  content: string;
  date?: string;
  relevance?: number;
}

export interface Verification {
  _id: string;
  statement_id: string;
  result: string;
  confidence?: number;
  verdict?: string;
  evidence?: Evidence[];
  notes?: string;
  metadata: {
    created_at: string;
    updated_at: string;
    verification_duration_ms?: number;
  };
  timeline: {
    checked_at: string;
    verified_by: string;
  };
}

export interface TimelineEvent {
  id: string;
  official_id: string;
  statement_id?: string;
  verification_id?: string;
  event_type: string;
  title: string;
  summary: string;
  occurred_at: string;
  result?: string;
}

export interface ApiListResponse<T> {
  items: T[];
}