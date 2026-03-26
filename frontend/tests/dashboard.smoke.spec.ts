import { expect, test } from '@playwright/test';

test('dashboard renders live backend data', async ({ page }) => {
	await page.goto('/');

	await expect(page).toHaveTitle(/Diputat Research Dashboard/i);
	await expect(page.getByRole('heading', { name: /Research dashboard for structured fact-checking/i })).toBeVisible();
	await expect(page.getByText('Loading dashboard data...')).not.toBeVisible();

	await expect(page.getByText('Kharkiv Reconstruction Update')).toBeVisible();
	await expect(page.getByText('Verification completed')).toBeVisible();

	await expect(page.locator('.metric-card').filter({ hasText: 'Officials' }).getByText('0')).toBeVisible();
	await expect(page.locator('.metric-card').filter({ hasText: 'Statements' }).getByText('20')).toBeVisible();
	await expect(page.locator('.metric-card').filter({ hasText: 'Verified statements' }).getByText('8')).toBeVisible();
	await expect(page.locator('.metric-card').filter({ hasText: 'Timeline events' }).getByText('28')).toBeVisible();
});

test('dashboard keeps only all-officials option when officials list is empty', async ({ page }) => {
	await page.goto('/');

	const options = page.locator('select option');
	await expect(options).toHaveCount(1);
	await expect(page.getByLabel('Focus official')).toHaveValue('all');
});