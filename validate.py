#!/usr/bin/env python3
"""Validate public template files and exact index hashes; no account/network access."""
import json,hashlib,re
from pathlib import Path
root=Path(__file__).parent
catalog=json.loads((root/'index.json').read_text());assert catalog['catalogVersion']==1 and catalog['repository']=='dingding229/mediatool-site'
ids=set()
for entry in catalog['templates']:
 assert re.fullmatch(r'[a-z0-9_]{1,80}',entry['id']) and entry['id']not in ids;ids.add(entry['id'])
 assert entry['file']=='sites/'+entry['id']+'.json' and entry['templateVersion']==1
 raw=(root/entry['file']).read_bytes();assert hashlib.sha256(raw).hexdigest()==entry['sha256']
 template=json.loads(raw);assert template['templateVersion']==1 and template['id']==entry['id'] and template['name']==entry['name']
 assert 'reuse_schema'not in template and len(template['requests'])==4
 for request in template['requests'].values():assert 'disabled_extends'not in request
assert {p.stem for p in(root/'sites').glob('*.json')}==ids
print('Validated',len(ids),'native templates; hashes matched, no account access.')
