import { useEffect, useState } from 'react';

interface Official {
  id: string;
  name: string;
  current_role: string;
  verification_score: number;
}

interface Statement {
  id: string;
  official_id: string;
  text: string;
  status: 'InProgress' | 'Done' | 'Blocked' | 'Toxic';
  source_url: string;
}

const OFFICIALS_URL = 'http://localhost:8080/api/v1/officials';
const STATEMENTS_URL = 'http://localhost:8080/api/v1/statements';
const FRIENDLY_BACKEND_ERROR = 'Backend connection issue. If this is a CORS error, we will patch Go next.';

function scoreBadgeClass(score: number): string {
  if (score < 30) {
    return 'bg-red-100 text-red-700 border-red-200';
  }

  if (score > 70) {
    return 'bg-green-100 text-green-700 border-green-200';
  }

  return 'bg-slate-100 text-slate-700 border-slate-200';
}

export default function App() {
  const [officials, setOfficials] = useState<Official[]>([]);
  const [statements, setStatements] = useState<Statement[]>([]);
  const [isOfficialsLoading, setIsOfficialsLoading] = useState<boolean>(true);
  const [isStatementsLoading, setIsStatementsLoading] = useState<boolean>(true);
  const [errorMessage, setErrorMessage] = useState<string>('');

  useEffect(() => {
    let isActive = true;

    async function loadOfficials(): Promise<void> {
      setIsOfficialsLoading(true);
      setErrorMessage('');

      try {
        const response = await fetch(OFFICIALS_URL);

        if (!response.ok) {
          throw new Error(`Request failed with status ${response.status}`);
        }

        const data: unknown = await response.json();
        const items = Array.isArray(data) ? data : (data as { items?: unknown }).items;

        if (!Array.isArray(items)) {
          throw new Error('Unexpected officials payload');
        }

        const parsedOfficials = items as Official[];

        if (isActive) {
          setOfficials(parsedOfficials);
        }
      } catch {
        if (isActive) {
          setOfficials([]);
          setErrorMessage(FRIENDLY_BACKEND_ERROR);
        }
      } finally {
        if (isActive) {
          setIsOfficialsLoading(false);
        }
      }
    }

    void loadOfficials();

    return () => {
      isActive = false;
    };
  }, []);

  useEffect(() => {
    let isActive = true;

    async function loadStatements(): Promise<void> {
      setIsStatementsLoading(true);

      try {
        const response = await fetch(STATEMENTS_URL);

        if (!response.ok) {
          throw new Error(`Request failed with status ${response.status}`);
        }

        const data: unknown = await response.json();
        const items = Array.isArray(data) ? data : (data as { items?: unknown }).items;

        if (!Array.isArray(items)) {
          throw new Error('Unexpected statements payload');
        }

        const parsedStatements = items as Statement[];

        if (isActive) {
          setStatements(parsedStatements);
        }
      } catch {
        if (isActive) {
          setStatements([]);
          setErrorMessage(FRIENDLY_BACKEND_ERROR);
        }
      } finally {
        if (isActive) {
          setIsStatementsLoading(false);
        }
      }
    }

    void loadStatements();

    return () => {
      isActive = false;
    };
  }, []);

  const inProgressStatements = statements.filter((statement) => statement.status === 'InProgress');
  const doneStatements = statements.filter((statement) => statement.status === 'Done');
  const blockedStatements = statements.filter((statement) => statement.status === 'Blocked');
  const toxicStatements = statements.filter((statement) => statement.status === 'Toxic');

  return (
    <div className="min-h-screen bg-slate-50 px-6 py-10 text-slate-900">
      <div className="mx-auto max-w-7xl">
        <header className="mb-8 rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
          <h1 className="text-3xl font-extrabold tracking-tight md:text-4xl">JIRA for the slaves in power</h1>
        </header>

        {isOfficialsLoading || isStatementsLoading ? (
          <p className="mb-6 text-sm text-slate-600">Loading dashboard data...</p>
        ) : null}

        {errorMessage ? (
          <div className="mb-6 rounded-xl border border-amber-200 bg-amber-50 p-4 text-amber-800">{errorMessage}</div>
        ) : null}

        {!errorMessage ? (
          <>
            <section className="grid grid-cols-1 gap-5 sm:grid-cols-2 xl:grid-cols-3">
              {officials.map((official) => (
                <article
                  key={official.id}
                  className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm transition hover:-translate-y-0.5 hover:shadow-md"
                >
                  <div className="flex items-start justify-between gap-4">
                    <div>
                      <h2 className="text-lg font-semibold text-slate-900">{official.name}</h2>
                      <p className="mt-1 text-sm text-slate-600">{official.current_role}</p>
                    </div>
                    <span
                      className={`inline-flex rounded-full border px-3 py-1 text-xs font-semibold ${scoreBadgeClass(official.verification_score)}`}
                    >
                      {official.verification_score}
                    </span>
                  </div>
                </article>
              ))}
            </section>

            <section className="mt-10">
              <div className="mb-4 flex items-center justify-between">
                <h2 className="text-2xl font-bold tracking-tight text-slate-900">Policy Execution Board</h2>
              </div>

              <div className="grid grid-cols-1 gap-4 lg:grid-cols-2 2xl:grid-cols-4">
                <KanbanColumn title="Backlog / In Progress" items={inProgressStatements} />
                <KanbanColumn title="Done / Kept Promises" items={doneStatements} />
                <KanbanColumn title="Blocked / Slacking" items={blockedStatements} />
                <KanbanColumn title="Toxic / Corrupt Action" items={toxicStatements} />
              </div>
            </section>
          </>
        ) : null}
      </div>
    </div>
  );
}

function KanbanColumn({ title, items }: { title: string; items: Statement[] }) {
  return (
    <section className="rounded-2xl border border-slate-300 bg-slate-200/60 p-3">
      <div className="mb-3 flex items-center justify-between">
        <h3 className="text-sm font-bold uppercase tracking-wide text-slate-700">{title}</h3>
        <span className="rounded-full bg-white px-2 py-1 text-xs font-semibold text-slate-600">{items.length}</span>
      </div>

      <div className="space-y-3">
        {items.map((statement) => (
          <article key={statement.id} className="rounded-xl border border-slate-200 bg-white p-4 shadow-sm">
            <p className="text-sm leading-6 text-slate-800">{statement.text}</p>
            <p className="mt-3 text-xs font-medium text-slate-500">Assigned to: Official #{statement.official_id}</p>
            <a
              href={statement.source_url}
              target="_blank"
              rel="noreferrer"
              className="mt-3 inline-flex items-center rounded-lg border border-slate-300 bg-slate-100 px-3 py-1.5 text-xs font-semibold text-slate-700 transition hover:bg-slate-200"
            >
              View Source
            </a>
          </article>
        ))}
        {items.length === 0 ? (
          <div className="rounded-xl border border-dashed border-slate-300 bg-slate-100 p-4 text-xs text-slate-500">
            No statements in this lane.
          </div>
        ) : null}
      </div>
    </section>
  );
}