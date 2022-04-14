DROP INDEX tokenapproval_protocolid;
DROP INDEX tokenpool_protocolid;

ALTER TABLE tokenapproval ADD COLUMN subject VARCHAR(1024);
UPDATE tokenapproval SET subject = protocol_id;

ALTER TABLE tokenpool RENAME COLUMN protocol_id TO locator;

CREATE UNIQUE INDEX tokenapproval_protocolid ON tokenapproval(connector, protocol_id);
CREATE UNIQUE INDEX tokenapproval_subject ON tokenapproval(pool_id, subject);
CREATE UNIQUE INDEX tokenpool_locator ON tokenpool(connector, locator);
