## 接口更改：
1. 序列化结构：
```go
    更改前：map[string]interface

    更改后：type MarshalInfo struct {
                // name  表示的是Campaign的字段名
                Name          string         `json:"name,omitempty"`
                // 表示的是在在query中传进来的condition的值
                QueryValue    interface{}    `json:"query_value,omitempty"`
                // 表示的是对应的id在索引中对应字段的值
                IndexValue    interface{}    `json:"index_value,omitempty"`
                // 操作符 eg: AndQuery: and  OrQuery:or, TermQuery: =, AndCheck: and_check
                Operation     string         `json:"operation,omitempty"`
                // check中的操作符，表示=,>,<等
                Op            operation.OP   `json:"op,omitempty"`
                // check中字段
                Transfer      bool           `json:"transfer,omitempty"`
                // 是否传入了自定义的operation对象
                SelfOperation bool           `json:"self_operation,omitempty"`
                // query中设置的label子弹
                Label         string         `json:"label,omitempty"`
                // 当前这个query的结果, 召回为true
                Result        bool           `json:"result"`

                Nodes         []*MarshalInfo `json:"nodes,omitempty"`
            }
```
2. 序列化接口:
  - Marshal() map[string]interface{}  ->  MarshalV2() *MarshalInfo
  - Unmarshal(idx index, info map[string]interface) Query -> UnmarshalV2(idx index, info *MarshalInfo) Query
3. search
  - 结构修改:
  ```go
          更改前: type Searcher struct {
                        Docs       []document.DocId
                        Time       time.Duration
                        FilterInfo map[document.DocId]map[string]interface{}
                        IndexDebug *debug.Debug
                        QueryDebug *debug.Debug
                  }
          更改后: type SearcherResult struct {
                        Docs []document.DocId
                        Time time.Duration
                 }
  ```
  - Search接口:
    ```
        将对象调用的方式改为函数的方式直接返回，不需要NewSearch
        func (s *Searcher) Search(iIndexer index.Index, query query.Query)
        func Search(iIndexer index.Index, query query.Query) *SearcherResult
    ```
  - Debug接口
    ```
        将对象调用的方式改为函数的方式直接返回，不需要NewSearch，同时将函数名改为Reply
        func (s *Searcher) Debug(idx index.Index, q map[string]interface{}, ids []document.DocId)
        func Replay(idx index.Index, info *marshal.MarshalInfo, ids []document.DocId) map[document.DocId]*marshal.MarshalInfo
    ```