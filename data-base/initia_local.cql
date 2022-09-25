CREATE KEYSPACE IF NOT EXISTS tech_radar
    WITH REPLICATION = {
        'class' : 'SimpleStrategy',
        'replication_factor' : 1
        };

CREATE TYPE IF NOT EXISTS tech_radar.technology_history(
    score tinyint,
    previous_score tinyint,
    added_date timestamp,
    comments text
);

CREATE TABLE IF NOT EXISTS tech_radar.technology(
    title text,
    friendly_title text,
    description text,
    team text static,
    moved tinyint,
    score tinyint,
    quadrant tinyint,
    active boolean,
    created_at timestamp,
    updated_at timestamp,
    history list<frozen<technology_history>>,
    PRIMARY KEY (team, friendly_title)
);

CREATE TABLE IF NOT EXISTS tech_radar.quadrant(
    id tinyint PRIMARY KEY,
    title text
)