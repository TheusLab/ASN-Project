from typing import List, Dict
from elasticsearch import Elasticsearch

es = Elasticsearch()

def search(query: str) -> List[Dict]:
    res = es.search(index="asn", body={"query": {"query_string": {"query": query}}})
    results = []
    for hit in res['hits']['hits']:
        results.append(hit["_source"])
    return results