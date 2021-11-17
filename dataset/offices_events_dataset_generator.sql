CREATE OR REPLACE FUNCTION random_between(min INT ,max INT)
    RETURNS INT AS
$$
BEGIN
    RETURN floor(random()* (max-min + 1) + min);
END;
$$ language 'plpgsql' STRICT;


CREATE OR REPLACE FUNCTION generate_payload(office_id INT)
    RETURNS jsonb AS
$$
BEGIN
    RETURN jsonb_build_object('id', office_id, 'name', 'test', 'description', 'test description', 'removed', false,  'created_at', now());
END;
$$ language 'plpgsql' STRICT;


INSERT INTO public.offices_events (office_id, type, status, payload, created_at)
SELECT series_id as office_id, random_between(1,5), 1, generate_payload(series_id) , now() FROM generate_series(1, 100000) as series_id;


