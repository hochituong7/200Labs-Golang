# enviroment
FROM ruby:2.7.1

RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs

# Install Yarn
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update -qq && apt-get install -y --no-install-recommends yarn
RUN gem install bundler -v 2.1.4
RUN mkdir /todos-app
WORKDIR /todos-app

ADD Gemfile /todos-app/Gemfile
ADD Gemfile.lock /todos-app/Gemfile.lock
RUN bundle install

ADD . /todos-app



source 'https://rubygems.org'
gem 'rails', '~> 6.0.5'


version: '2'
services:
  app:
    build: .
    command: bundle exec rails s -version: '2'
services:
  app:
    build: .
    command: bundle exec rails s -p 3003 -b '0.0.0.0'
    volumes:
      - .:/todos-app
    ports:
      - "3000:3003"
  webserver:
    image: nginx:alpine
    container_name: webserver
    restart: unless-stopped
    tty: true
    ports:
      - "80:80"
      - "443:443"
    networks:
      - app-network
  sqlite3:
    image: nouchka/sqlite3:latest
    stdin_open: true
    tty: true
    volumes:
      - ./db/:/root/db/
    ports:
      - '9000:9000'
            
networks:
  app-network:
    driver: bridge
  

