import { useEffect, useMemo, useState } from 'react';
import { fetchOfficials, fetchStatements, fetchTimeline, fetchVerifications } from './api';
import type { Official, Statement, TimelineEvent, Verification } from './models';

type LoadState = 'idle' | 'loading' | 'ready' | 'error';

function formatDate(value: string): string {
  return new Date(value).toLocaleString();
}

function resultLabel(result?: string): string {
  switch (result) {
    case 'true':
      return 'True';
    case 'false':
      return 'False';
    case 'partially':
      return 'Partially true';
    case 'misleading':
      return 'Misleading';
    case 'unverifiable':
      return 'Unverifiable';
    default:
      return 'Pending';
  }
}

export default function App() {
  const [state, setState] = useState<LoadState>('idle');
  const [error, setError] = useState<string>('');
  const [selectedOfficialId, setSelectedOfficialId] = useState<string>('all');
  const [officials, setOfficials] = useState<Official[]>([]);
  const [statements, setStatements] = useState<Statement[]>([]);
  const [verifications, setVerifications] = useState<Verification[]>([]);
  const [timeline, setTimeline] = useState<TimelineEvent[]>([]);

  useEffect(() => {
    let active = true;

    async function loadDashboard() {
      setState('loading');
      setError('');

      try {
        const officialFilter = selectedOfficialId === 'all' ? undefined : selectedOfficialId;
        const [nextOfficials, nextStatements, nextVerifications, nextTimeline] = await Promise.all([
          fetchOfficials(),
          fetchStatements(officialFilter),
          fetchVerifications(),
          fetchTimeline(officialFilter)
        ]);

        if (!active) {
          return;
        }

        setOfficials(nextOfficials);
        setStatements(nextStatements);
        setVerifications(nextVerifications);
        setTimeline(nextTimeline);
        setState('ready');
      } catch (loadError) {
        if (!active) {
          return;
        }

        setError(loadError instanceof Error ? loadError.message : 'Unknown error');
        setState('error');
      }
    }

    void loadDashboard();

    return () => {
      active = false;
    };
  }, [selectedOfficialId]);

  const verificationsByStatementId = useMemo(() => {
    return new Map(verifications.map((verification) => [verification.statement_id, verification]));
  }, [verifications]);

  const statementCounts = useMemo(() => {
    const verified = statements.filter((statement) => statement.status === 'verified').length;

    return {
      officials: officials.length,
      statements: statements.length,
      verified,
      timeline: timeline.length
    };
  }, [officials.length, statements, timeline.length]);

  return (
    <div className="page-shell">
      <header className="hero">
        <div>
          <p className="eyebrow">Diputat</p>
          <h1>Research dashboard for structured fact-checking</h1>
          <p className="hero-copy">
            The MVP keeps the workflow deterministic: statements, evidence, and verdicts are loaded from local JSON data through the Go API.
          </p>
        </div>
        <label className="filter-panel">
          <span>Focus official</span>
          <select value={selectedOfficialId} onChange={(event) => setSelectedOfficialId(event.target.value)}>
            <option value="all">All officials</option>
            {officials.map((official) => (
              <option key={official._id} value={official._id}>
                {official.name}
              </option>
            ))}
          </select>
        </label>
      </header>

      <section className="metrics-grid">
        <article className="metric-card">
          <span>Officials</span>
          <strong>{statementCounts.officials}</strong>
        </article>
        <article className="metric-card">
          <span>Statements</span>
          <strong>{statementCounts.statements}</strong>
        </article>
        <article className="metric-card">
          <span>Verified statements</span>
          <strong>{statementCounts.verified}</strong>
        </article>
        <article className="metric-card">
          <span>Timeline events</span>
          <strong>{statementCounts.timeline}</strong>
        </article>
      </section>

      {state === 'loading' ? <p className="status-banner">Loading dashboard data...</p> : null}
      {state === 'error' ? <p className="status-banner error">Unable to load data: {error}</p> : null}

      <main className="content-grid">
        <section className="panel">
          <div className="panel-header">
            <h2>Officials</h2>
            <span>{officials.length} tracked</span>
          </div>
          <div className="stack-list">
            {officials.map((official) => (
              <article className="stack-card" key={official._id}>
                <div>
                  <h3>{official.name}</h3>
                  <p>{official.position}</p>
                </div>
                <span>{official.party ?? 'Non-partisan'}</span>
              </article>
            ))}
          </div>
        </section>

        <section className="panel wide">
          <div className="panel-header">
            <h2>Statements</h2>
            <span>{statements.length} loaded</span>
          </div>
          <div className="stack-list">
            {statements.map((statement) => {
              const verification = verificationsByStatementId.get(statement._id);

              return (
                <article className="statement-card" key={statement._id}>
                  <div className="statement-meta-row">
                    <span>{statement.source.media_outlet ?? statement.source.type}</span>
                    <span>{formatDate(statement.source.date)}</span>
                  </div>
                  <h3>{statement.source.title ?? 'Untitled source'}</h3>
                  <p>{statement.content}</p>
                  <div className="tag-row">
                    <span className="tag">{statement.status}</span>
                    {statement.topics?.map((topic) => (
                      <span className="tag muted" key={topic}>
                        {topic}
                      </span>
                    ))}
                    <span className={`tag verdict ${verification?.result ?? 'pending'}`}>
                      {resultLabel(verification?.result)}
                    </span>
                  </div>
                </article>
              );
            })}
          </div>
        </section>

        <section className="panel">
          <div className="panel-header">
            <h2>Verifications</h2>
            <span>{verifications.length} completed</span>
          </div>
          <div className="stack-list">
            {verifications.map((verification) => (
              <article className="stack-card verification-card" key={verification._id}>
                <div>
                  <h3>{resultLabel(verification.result)}</h3>
                  <p>{verification.verdict}</p>
                </div>
                <span>{Math.round((verification.confidence ?? 0) * 100)}%</span>
              </article>
            ))}
          </div>
        </section>

        <section className="panel wide">
          <div className="panel-header">
            <h2>Timeline</h2>
            <span>{timeline.length} events</span>
          </div>
          <div className="timeline-list">
            {timeline.map((event) => (
              <article className="timeline-item" key={event.id}>
                <div className="timeline-dot" />
                <div>
                  <div className="statement-meta-row">
                    <span>{event.title}</span>
                    <span>{formatDate(event.occurred_at)}</span>
                  </div>
                  <p>{event.summary}</p>
                  {event.result ? <span className={`tag verdict ${event.result}`}>{resultLabel(event.result)}</span> : null}
                </div>
              </article>
            ))}
          </div>
        </section>
      </main>
    </div>
  );
}