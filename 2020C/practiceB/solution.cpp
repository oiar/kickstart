#include<bits/stdc++.h>
using namespace std;

const int N = 26;
unordered_set<int> dependOn[N];
unordered_set<int> base[N];
bool used[N];
string answer;

// b将自己所依赖字母d剔除后，不依赖其他字母，则记录下来并继续 dfs b被依赖的字母
void dfs(int d) {
  answer += d + 'A';
  used[d] = false;
  for(auto b : base[d]) {
    dependOn[b].erase(d);
    if(used[b] && dependOn[b].size() == 0)
      dfs(b);
  }
}

void init() {
  for(int i = 0; i < N; i++) {
    dependOn[i].clear();
    base[i].clear();
  }
  memset(used, false, sizeof(used));
  answer = "";
}

int main() {
  int T, output = 0;
  cin >> T;
  while (T--)
  {
    init();
    int R, C, count = 0;
    cin >> R >> C;
    char letter[35][35];
    for(int i = 0; i < R; i ++)
      cin >> letter[i];
    for(int i = 0; i < R - 1; i++)
      for(int j = 0; j < C; j++)
        if(letter[i][j] != letter[i + 1][j]) {
          int u = letter[i][j] - 'A', d = letter[i + 1][j] - 'A';
          // 维护一个依赖数组和一个被依赖数组
          dependOn[u].insert(d);
          base[d].insert(u);
          // 维护26个字母中被使用的字母
          used[u] = true;
          used[d] = true;
        }
    for(int i = 0; i < C; i++)
      used[letter[R-1][i] - 'A'] = true;

    // 统计出现的不同字母个数
    for(int i = 0; i < N; i++)
      if(used[i]) count++;

    // 若当前字母没有依赖下层字母，dfs 依赖于当前字母的其他字母
    for(int i = 0; i < N; i++)
      if(used[i] && dependOn[i].size() == 0)
        dfs(i);
    printf("Case #%d: ", ++output);
    if(answer.length() != count)
      cout << "-1\n";
    else cout << answer << endl;
  }
  
  return 0;
}