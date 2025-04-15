-- Trigger function to update `updated_date` whenever an article is updated
CREATE OR REPLACE FUNCTION update_updated_date_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_date = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger that will fire after every update to the articles table
CREATE TRIGGER update_article_updated_date
AFTER UPDATE ON articles
FOR EACH ROW
EXECUTE FUNCTION update_updated_date_column();
