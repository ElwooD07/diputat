import type {
  ApiListResponse,
  Official,
  Statement,
  TimelineEvent,
  Verification
} from './models';

const apiBaseUrl = import.meta.env.VITE_API_URL ?? 'http://localhost:8080/api/v1';

async function getList<T>(path: string): Promise<T[]> {
  const response = await fetch(`${apiBaseUrl}${path}`);
  if (!response.ok) {
    throw new Error(`Request failed with status ${response.status}`);
  }

  const payload = (await response.json()) as ApiListResponse<T>;
  return payload.items;
}

export function fetchOfficials(): Promise<Official[]> {
  return getList<Official>('/officials');
}

export function fetchStatements(officialId?: string): Promise<Statement[]> {
  const query = officialId ? `?official_id=${encodeURIComponent(officialId)}` : '';
  return getList<Statement>(`/statements${query}`);
}

export function fetchVerifications(): Promise<Verification[]> {
  return getList<Verification>('/verifications');
}

export function fetchTimeline(officialId?: string): Promise<TimelineEvent[]> {
  const query = officialId ? `?official_id=${encodeURIComponent(officialId)}` : '';
  return getList<TimelineEvent>(`/timeline${query}`);
}