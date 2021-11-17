INSERT INTO public.offices ("name", description, removed, created_at)
SELECT 'test', 'test description', false, now() FROM generate_series(1, 100000);
