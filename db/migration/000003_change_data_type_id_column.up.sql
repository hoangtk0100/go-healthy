ALTER TABLE body_records ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE meals ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE exercises ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE diaries ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE blog_posts ALTER COLUMN id SET DEFAULT gen_random_uuid();
