//
//    Copyright 2020 Satvik Reddy
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package main

import (
	"github.com/SatvikR/cyandb/server"
	"log"
)

func main() {
	db := server.CreateServer(server.DefaultDBPath)

	/*	_, _ = db.Set("hello", "world")
		_, _ = db.Set("chess", "is cool")
		_, _ = db.Set("chocolate", "cake")
	*/

	if _, err := db.Set("chocolate", "cake!"); err != nil {
		log.Fatal(err)
	}

}
