#include <bits/stdc++.h>
using namespace std;
#define N 55
#define LL long long
vector<pair<int,int>> g[N];
bool vis[N];

void dfs(int node) {
  vis[node] = true;
  for(int i = 0; i < g[node].size(); i++) {
    int next = g[node][i].first;
    if(!vis[next]) {
      dfs(next);
    }
  }
}

void update(int a, int b, int c) {
  if(g[a].size() < 1) {
    g[a].push_back(make_pair(b, c));
  } else if(g[a][0].second > c) {
    g[a][0] = make_pair(b, c);
  }
}

int main() {
  int T, o = 0;
  cin >> T;
  while (T--)
  {
    int node, edge = 0;
    cin >> node >> edge;
    for(int i = 0; i < node; i++) {
      g[i].clear();
    }
    for(int i = 0; i < edge; i++) {
      int a, b, c;
      cin >> a >> b >> c;
      a--; b--;
      update(a, b, c);
      update(b, a, c);
    }
    unordered_set<int> zeros;
    for(int i = 0; i < node; i++) {
      if(g[i].size()) {
        int b = g[i][0].first;
        int c = g[i][0].second;
        if (g[i][0].second != g[b][0].second) {
          g[b].push_back(make_pair(i, c));
        }
        if(c == 0) {
          zeros.insert(i);
          zeros.insert(b);
        }
      }
    }

    int part = 0;
    memset(vis, false, sizeof(vis));
    for(int i = 0; i < node; i++) {
      if(!vis[i]) {
        dfs(i);
        part++;
      }
    }

    int to_zeros = 0;
    if(zeros.size()) {
      for(int i = 0; i < node; i++) {
        if(g[i].size()) {
          int a = g[i][0].first;
          if(zeros.find(i) == zeros.end() && zeros.find(a) != zeros.end()) {
            to_zeros++;
          }
        }
      }
    }
    LL res = 1LL << (part + to_zeros);
    printf("Case #%d: %lld\n", ++o, res);
  }
  
  return 0;
}