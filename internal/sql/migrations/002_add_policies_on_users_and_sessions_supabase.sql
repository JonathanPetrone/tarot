-- +goose Up
ALTER TABLE public.users ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.sessions ENABLE ROW LEVEL SECURITY;

CREATE POLICY "users_own_data" ON public.users
FOR ALL USING (auth.uid()::text = id::text);

CREATE POLICY "sessions_own_data" ON public.sessions  
FOR ALL USING (auth.uid()::text = user_id::text);

-- +goose Down
DROP POLICY IF EXISTS "users_own_data" ON public.users;
DROP POLICY IF EXISTS "sessions_own_data" ON public.sessions;
ALTER TABLE public.users DISABLE ROW LEVEL SECURITY;
ALTER TABLE public.sessions DISABLE ROW LEVEL SECURITY;