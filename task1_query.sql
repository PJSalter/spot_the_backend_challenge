SELECT
-- Selects the name column representing the spot name
  name,
  -- Extracts the domain from the website using regular expression
  REGEXP_REPLACE(website, '(https?://)([^/]+)(.*)', '\2') AS domain,
  -- Counts the number of occurrences for each spot name and domain combination
  COUNT(*) AS count_number
FROM
  spots
GROUP BY
-- Groups the data by spot name and domain
  name, REGEXP_REPLACE(website, '(https?://)([^/]+)(.*)', '\2')
HAVING
-- Filters for domains with a count greater than 1
  COUNT(*) > 1;

