/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cache

import (
	"net/http"
	"strings"
)

type Cache interface {
	Get(key string) *http.Response
	Set(key string, value *http.Response)
	Delete(key string)
	Clear()
	Has(key string) bool
}

func CreateCacheKey(req *http.Request) string {
	s := req.URL.Host + req.URL.Path
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, ".", "_", -1)
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, "@", "_", -1)
	s = strings.Replace(s, "+", "_", -1)
	s = strings.Replace(s, ":", "_", -1)
	return s
}