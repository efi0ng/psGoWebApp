package models

import ()

type Category struct {
	imageUrl      string
	title         string
	description   string
	id            int
}

func (this *Category) ImageUrl() string {
	return this.imageUrl
}

func (this *Category) Title() string {
	return this.title
}

func (this *Category) Description() string {
	return this.description
}

func (this *Category) Id() int {
	return this.id
}

func (this *Category) SetImageUrl(value string) {
	this.imageUrl = value
}

func (this *Category) SetTitle(value string) {
	this.title = value
}

func (this *Category) SetDescription(value string) {
	this.description = value
}

func (this *Category) SetId(value int) {
	this.id = value
}

func GetCategories() []Category {
	result := []Category{
		Category{
			id:       1,
			imageUrl: "lemon.png",
			title:    "Juices and Mixes",
			description: `Explore our wide assortment of juices and mixes expected by
							today's lemonade stand clientelle. Now featuring a full line of
							organic juices that are guaranteed to be obtained from trees that
							have never been treated with pesticides or artificial
							fertilizers.`,
		},
		Category{
			id:       2,
			imageUrl: "kiwi.png",
			title:    "Cups, Straws, and Other Supplies",
			description: `From paper cups to bio-degradable plastic to straws and
						napkins, LSS is your source for the sundries that keep your stand
						running smoothly.`,
		},
		Category{
			id:       3,
			imageUrl: "pineapple.png",
			title:    "Signs and Advertising",
			description: `Sure, you could just wait for people to find your stand
						along the side of the road, but if you want to take it to the next
						level, our premium line of advertising supplies.`,
		}}
	return result
}
