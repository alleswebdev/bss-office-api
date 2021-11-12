INSERT INTO public.offices ("name", description, removed, created)
SELECT 'test', 'test description', false, now() FROM generate_series(1, 100000);
