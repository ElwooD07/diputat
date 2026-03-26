import { defineConfig } from '@playwright/test';

export default defineConfig({
	testDir: './tests',
	fullyParallel: false,
	timeout: 30_000,
	retries: 0,
	use: {
		baseURL: 'http://127.0.0.1:4200',
		headless: true,
		trace: 'retain-on-failure'
	},
	webServer: [
		{
			command: 'export PATH="$HOME/go/bin:$PATH" && go run ./cmd/server',
			cwd: '../backend',
			url: 'http://127.0.0.1:8080/health',
			reuseExistingServer: true,
			timeout: 60_000
		},
		{
			command: 'npm run dev -- --host 127.0.0.1 --port 4200',
			cwd: '.',
			url: 'http://127.0.0.1:4200',
			reuseExistingServer: true,
			timeout: 60_000
		}
	]
});