import { useEffect, useState } from "react";

type DependencyStatus = "ok" | "down";

type HealthResponse = {
  status: "ok" | "degraded";
  service: string;
  version: string;
  timestamp: string;
  dependencies: Record<string, DependencyStatus>;
};

async function fetchHealth(): Promise<HealthResponse> {
  const response = await fetch("/api/v1/health");
  if (!response.ok) {
    throw new Error(`Health request failed with status ${response.status}`);
  }
  return response.json();
}

function App() {
  const [health, setHealth] = useState<HealthResponse | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    fetchHealth()
      .then((data) => {
        if (active) {
          setHealth(data);
        }
      })
      .catch((requestError: Error) => {
        if (active) {
          setError(requestError.message);
        }
      })
      .finally(() => {
        if (active) {
          setLoading(false);
        }
      });

    return () => {
      active = false;
    };
  }, []);

  const connectionLabel =
    loading ? "检查中" : health ? "已连接" : "未连接";
  const connectionClass = loading
    ? "neutral"
    : health
      ? health.status === "ok"
        ? "healthy"
        : "degraded"
      : "offline";

  return (
    <div className="app">
      <header className="topbar">
        <h1>StockAI</h1>
        <span className="stage">V0.1</span>
      </header>

      <main className="content">
        <section className="panel">
          <div className="panel-header">
            <h2>Backend Connection</h2>
            <span className={`status ${connectionClass}`}>{connectionLabel}</span>
          </div>

          {loading && <p className="message">正在请求后端健康检查接口。</p>}
          {!loading && error && <p className="message error">{error}</p>}
          {!loading && health && (
            <dl className="details">
              <div>
                <dt>Service</dt>
                <dd>{health.service}</dd>
              </div>
              <div>
                <dt>Version</dt>
                <dd>{health.version}</dd>
              </div>
              <div>
                <dt>Updated</dt>
                <dd>{new Date(health.timestamp).toLocaleString("zh-CN")}</dd>
              </div>
              <div>
                <dt>Dependencies</dt>
                <dd>
                  {Object.entries(health.dependencies)
                    .map(([name, status]) => `${name}: ${status}`)
                    .join(", ")}
                </dd>
              </div>
            </dl>
          )}
        </section>
      </main>
    </div>
  );
}

export default App;

