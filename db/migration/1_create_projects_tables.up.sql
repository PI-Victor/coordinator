CREATE TABLE PROJECTS (
  ID   SERIAL PRIMARY KEY,
  NAME VARCHAR(256) NOT NULL UNIQUE,
  TYPE VARCHAR(256) NOT NULL
);

CREATE TABLE REPOSITORIES (
  ID         SERIAL PRIMARY KEY,
  NAME       VARCHAR(256) NOT NULL,
  PROJECT_ID SERIAL REFERENCES PROJECTS (ID) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE BRANCHES (
  ID            SERIAL PRIMARY KEY,
  NAME          VARCHAR(256) NOT NULL,
  REPOSITORY_ID SERIAL REFERENCES REPOSITORIES (ID) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE BRANCH_META (
  ID        SERIAL PRIMARY KEY,
  BRANCH_ID INT REFERENCES BRANCHES (ID) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE LANGUAGES (
  ID         SERIAL PRIMARY KEY,
  NAME       VARCHAR(256),
  PERCENTAGE FLOAT,
  UNIQUE (ID, NAME)
);

CREATE TABLE BRANCH_LANGUAGES (
  LANGUAGE_ID    INT REFERENCES LANGUAGES (ID),
  BRANCH_META_ID INT REFERENCES BRANCH_META (ID),
  UNIQUE (LANGUAGE_ID, BRANCH_META_ID)
);

CREATE TABLE COMMITS (
  ID           SERIAL,
  SHA          VARCHAR(256),
  MESSAGE      VARCHAR(256),
  TIMESTAMP    TIMESTAMP NULL,
  AUTHOR_NAME  VARCHAR(256),
  AUTHOR_EMAIL VARCHAR(256),
  ADDED        VARCHAR(256) [],
  MODIFIED     VARCHAR(256) [],
  REMOVED      VARCHAR(256) [],
  BRANCH_ID    INT REFERENCES BRANCHES (ID) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (ID, SHA)
);