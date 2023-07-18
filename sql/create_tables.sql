create table gin_sample.weather_points (
  id int unique not null auto_increment,
  latitude decimal(18, 16),
  longitude decimal(19, 16),
  timezone varchar(256),
  created_at datetime,
  updated_at datetime
);

create table gin_sample.daily_weathers (
  id int unique not null auto_increment,
  weather_point_id int,
  weather_code int,
  max_temperature decimal(4, 2),
  min_temperature decimal(4, 2),
  weather_date date,
  created_at datetime,
  updated_at datetime
);

create table gin_sample.news_articles (
  id int unique not null auto_increment,
  title varchar(256),
  description text,
  content text,
  article_url varchar(1024),
  image_url varchar(1024),
  resource_name varchar(256),
  created_at datetime,
  updated_at datetime
);
