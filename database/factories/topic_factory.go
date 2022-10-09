package factories

import (
    "Gohub/app/models/topic"

    "github.com/bxcodec/faker/v3"
)

func MakeTopics(count int) []topic.Topic {

    var objs []topic.Topic

    // 设置唯一性，如 Topic 模型的某个字段需要唯一，即可取消注释
    // faker.SetGenerateUniqueValues(true)

    for i := 0; i < count; i++ {
        topicModel := topic.Topic{
            Title:        faker.Username(),
			Body:         faker.Sentence(),
            UserID:       "2",
            CategoryID:   "3",

        }
        objs = append(objs, topicModel)
    }

    return objs
}