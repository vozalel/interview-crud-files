create schema acl;

CREATE TABLE acl.user_datasource (
                                     user_id int4 NOT NULL,
                                     datasource_name varchar(100) NOT NULL,
                                     "read" bool NOT NULL,
                                     "update" bool NOT NULL,
                                     "delete" bool NOT NULL,
                                     "grant" bool NOT NULL,
                                     "revoke" bool NOT NULL,
                                     CONSTRAINT acl_un UNIQUE (user_id, datasource_name)
);


CREATE TABLE acl.user_perform (
                                  user_id int4 NOT NULL,
                                  "create" bool NOT NULL,
                                  CONSTRAINT user_perform_un UNIQUE (user_id)
);