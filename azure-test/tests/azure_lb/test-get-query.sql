select name, id, type, region, resource_group, subscription_id, tags
from azure.azure_lb
where name = '{{ resourceName }}' and resource_group = '{{ resourceName }}';
