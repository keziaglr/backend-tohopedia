package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/keziaglr/backend-tohopedia/graph"
	"github.com/keziaglr/backend-tohopedia/graph/generated"
	"github.com/keziaglr/backend-tohopedia/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

var db *gorm.DB;

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000", "http://localhost:8080"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB : db,
	}}))
			
	router.Handle("/", playground.Handler("Tohopedia", "/query"))
	router.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
    if err != nil {
        panic(err)
    }
}

func initDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/tohopedia?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err)
    }
	
	migrate();
	seeds();
}

func migrate(){
	
	db.Migrator().DropTable(
		&model.Campaign{},
		&model.ShippingVendor{},
		&model.Badges{},  
		&model.Voucher{},
		&model.Category{},
		&model.SubCategory{}, 
		&model.ProductImage{}, 
		&model.Product{},
		&model.User{}, 
		&model.Shop{}, 
		&model.ShopShippingVendor{},
		&model.ShopVoucher{},
		&model.UserVoucher{},
		&model.UserWishlist{}, 
		&model.Cart{}, 
		&model.Chat{}, 
		&model.ChatDetail{}, 
		&model.Review{}, 
		&model.ReviewReply{}, 
		&model.Discussion{}, 
		&model.DiscussionReply{}, 
		&model.Otp{}, 
		&model.Request{}, 
		&model.Review{},  
		&model.TransactionHeader{}, 
		&model.TransactionDetail{},
	)
	
	db.Exec("DROP TABLE product_images, shop_product, user_wishlist, user_voucher")
	db.AutoMigrate(
		&model.Campaign{},
		&model.ShippingVendor{},
		&model.Badges{},  
		&model.Voucher{},
		&model.Category{},
		&model.SubCategory{}, 
		&model.ProductImage{}, 
		&model.Product{},
		&model.User{}, 
		&model.Shop{},
		&model.ShopShippingVendor{},
		&model.ShopVoucher{},
		&model.UserVoucher{},
		&model.UserWishlist{}, 
		&model.Cart{}, 
		&model.Chat{}, 
		&model.ChatDetail{}, 
		&model.Review{}, 
		&model.ReviewReply{}, 
		&model.Discussion{}, 
		&model.DiscussionReply{}, 
		&model.Otp{}, 
		&model.Request{}, 
		&model.Review{},  
		&model.TransactionHeader{}, 
		&model.TransactionDetail{},
	)	
}

func seeds(){
	seedMaster();
}

func seedMaster() {
	//User
	user := []model.User{
		{
			Email: "kezia@mail.com",
			Password: "kezia123",
			Name: "kekez",
			Dob: "2021-03-03",
			Gender: "Female",
			PhoneNumber: "282892929",
			ProfilePicture: "https://i.mydramalist.com/vK4lp_5_c.jpg",
			IsSuspend: false,
			ShippingAddress: "Kalimantan Barat",
			Role: "User",
		}, {
			Email: "admin@admin.com",
			Password: "admin123",
			Name: "Admin",
			Dob: "2021-05-05",
			Gender: "Male",
			PhoneNumber: "2992929210",
			ProfilePicture: "https://awsimages.detik.net.id/community/media/visual/2018/02/15/f91bd7e4-25b5-4ac7-b1a1-4844d3a3b89b.jpeg?w=1200",
			IsSuspend: false,
			ShippingAddress: "Jakarta",
			Role: "Admin",
		}, {
			Email: "twice@mail.com",
			Password: "twice123",
			Name: "twice",
			Dob: "2021-03-03",
			Gender: "Female",
			PhoneNumber: "71028392012",
			ProfilePicture: "https://assets.pikiran-rakyat.com/crop/16x22:716x737/x/photo/2021/09/16/103691333.jpeg",
			IsSuspend: false,
			ShippingAddress: "Jawa Timur",
			Role: "User",
		}, {
			Email: "kep1er@mail.com",
			Password: "kep1er",
			Name: "kep1er",
			Dob: "2021-04-04",
			Gender: "Female",
			PhoneNumber: "038389234",
			ProfilePicture: "https://staticg.sportskeeda.com/editor/2022/01/b3203-16416582079625-1920.jpg",
			IsSuspend: false,
			ShippingAddress: "Depok",
			Role: "User",
		},
	}

	//Badge
	badge := []model.Badges{
		{
			StartPoint: 1,
			EndPoint: 50,
			Badge: "Bronze",
		},{
			StartPoint: 51,
			EndPoint: 100,
			Badge: "Silver",
		},{
			StartPoint: 101,
			EndPoint: 150,
			Badge: "Gold",
		}, {
			StartPoint: 151,
			EndPoint: 200,
			Badge: "Diamond",
		},
	}

	
	category := []model.Category{
		{
			Name: "Furniture",
		},{
			Name: "Electronic",
		},{
			Name: "Fashion",
		},{
			Name: "Health",
		},{
			Name: "Beauty",
		},{
			Name: "Office & Stationery",
		},{
			Name: "Food & Beverage",
		},{
			Name: "Sport",
		},
	}

	subCategory := []model.SubCategory{
		{
			CategoryID: 1,
			Name: "Chair",
		}, {
			CategoryID: 1,
			Name: "Table",
		},{
			CategoryID: 1,
			Name: "Bed",
		},{
			CategoryID: 2,
			Name: "Handphone",
		}, {
			CategoryID: 2,
			Name: "Camera",
		}, {
			CategoryID: 2,
			Name: "Laptop",
		}, {
			CategoryID: 3,
			Name: "Dress",
		}, {
			CategoryID: 3,
			Name: "Shirt",
		}, {
			CategoryID: 3,
			Name: "Jeans",
		}, {
			CategoryID: 4,
			Name: "Medicine",
		}, {
			CategoryID: 4,
			Name: "Vitamin",
		}, {
			CategoryID: 4,
			Name: "Essential Oil",
		}, {
			CategoryID: 5,
			Name: "Skincare",
		}, {
			CategoryID: 5,
			Name: "Hair Care",
		}, {
			CategoryID: 5,
			Name: "Nail Art",
		}, {
			CategoryID: 6,
			Name: "Stationery",
		}, {
			CategoryID: 6,
			Name: "Book",
		}, {
			CategoryID: 6,
			Name: "Paper",
		}, {
			CategoryID: 7,
			Name: "Snack",
		}, {
			CategoryID: 7,
			Name: "Cake",
		}, {
			CategoryID: 7,
			Name: "Frozen Food",
		}, {
			CategoryID: 8,
			Name: "Ball",
		}, {
			CategoryID: 8,
			Name: "Racket",
		}, {
			CategoryID: 8,
			Name: "Hiking & Camping",
		},
	}

	shipping := []model.ShippingVendor{
		{
			Name: "SiCepat",
			DeliveryTime: 5,
			Price: 10000,
		}, {
			Name: "JNE",
			DeliveryTime: 10,
			Price: 30000,
		}, {
			Name: "JNT",
			DeliveryTime: 20,
			Price: 50000,
		}, {
			Name: "AnterAja",
			DeliveryTime: 30,
			Price: 90000,
		}, {
			Name: "Grab Instant",
			DeliveryTime: 1,
			Price: 70000,
		},
	}

	voucher := []model.Voucher{
		{
			Name: "Voucher 1",
			Description: "Desc Voucher 1",
			DiscountRate: 50,
			Tnc: "TNC Voucher 1",
			StartTime: time.Now(),
			EndTime: time.Now().Add(5),
		},{
			Name: "Voucher 2",
			Description: "Desc Voucher 2",
			DiscountRate: 70,
			Tnc: "TNC Voucher 2",
			StartTime: time.Now(),
			EndTime: time.Now().Add(10),
		},{
			Name: "Voucher 3",
			Description: "Desc Voucher 3",
			DiscountRate: 30,
			Tnc: "TNC Voucher 3",
			StartTime: time.Now(),
			EndTime: time.Now().Add(15),
		},
	}
	
	shop := []model.Shop{
		{
			UserID: 1,
			Name: "iBox",
			NameSlug: "iBox Official",
			Points: 20,
			Image: "https://cf.shopee.co.id/file/eb88e2ccea222edb18618943dc4807ab",
			OperationalStatus: "Open",
			OperationalHour: "09.00 - 15.00",
			Description: "Desc iBox",
			Slogan: "Slogan iBox",
			Address: "Alamat iBox",
			PhoneNumber: "018282292",
			BadgesID: 1,
			Product: []*model.Product{
				{
					Name: "MacBook Pro",
					Description: "MacBook Pro Description",
					Price: 50000000,
					Discount: 10,
					MetaData: "Weight: 3gram",
					AddedTime: time.Now(),
					SoldCount: 100,
					Stock: 10000,
					Rating: 5,
					SubCategoryID:6,
					Images: []*model.ProductImage{
						{
							URL: "https://static.bmdstatic.com/pk/product/medium/600fc3d770293.jpg",
						},{
							URL: "https://cdn.medcom.id/images/content/2022/01/25/1381007/pKZzQ6uyCr.jpg",
						},
					},
				},{
					Name: "IPhone 13",
					Description: "IPhone 13 Description",
					Price: 700000,
					Discount: 2,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 200000,
					SoldCount: 250,
					Rating: 4,
					SubCategoryID:5,
					Images: []*model.ProductImage{
						{
							URL: "https://www.notebookcheck.net/fileadmin/Notebooks/News/_nc3/iphone_12_pro_iphone_12_iphone_11_battery_life.jpg",
						},{
							URL: "https://cdn.lumen.id/commerce/digimap/file/a4aa2367-2527-4a98-9b61-2db0468ba61c/PDP-iPhone-13-Pink-4-medium.jpeg",
						},
					},
				},{
					Name: "IPhone 12",
					Description: "IPhone 12 Description",
					Price: 600000,
					Discount: 3,
					SoldCount: 250,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 200000,
					Rating: 5,
					SubCategoryID:5,
					Images: []*model.ProductImage{
						{
							URL: "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full//99/MTA-8107982/apple_iphone_11_64gb_full01_ge72ewvv.jpg",
						},{
							URL: "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/refurb-iphone-11-pro-gold-2019?wid=2000&hei=2000&fmt=jpeg&qlt=80&.v=1611101491000",
						},
					},
				},{
					Name: "IPhone 11",
					Description: "IPhone 11 Description",
					Price: 650000,
					Discount: 4,
					SoldCount: 300,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 50000,
					Rating: 4,
					SubCategoryID:5,
					Images: []*model.ProductImage{
						{
							URL: "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full//96/MTA-3567047/apple_apple-iphone-x-64-gb-smartphone_full05.jpg",
						},{
							URL: "https://images.tokopedia.net/img/cache/500-square/product-1/2019/5/3/60479447/60479447_b59bd146-2fd3-4c6c-80fe-0fc18fae89fb_529_529",
						},
					},
				},{
					Name: "IPhone X",
					Description: "IPhone X Description",
					Price: 500000,
					Discount: 7,
					SoldCount: 400,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 70000,
					Rating: 5,
					SubCategoryID:5,
					Images: []*model.ProductImage{
						{
							URL: "https://id-live-05.slatic.net/original/35fd56ea9616f0353f925169c8950e96.jpg_720x720q80.jpg_.webp",
						},{
							URL: "https://id-live-05.slatic.net/p/ea1b93cbd237ecef42d4b90978134b38.jpg_720x720q80.jpg_.webp",
						},
					},
				},{
					Name: "IPhone 8",
					Description: "IPhone 8 Description",
					Price: 800000,
					Discount: 8,
					SoldCount: 600,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 80000,
					Rating: 4,
					SubCategoryID:5,
					Images: []*model.ProductImage{
						{
							URL: "https://cdn.alloallo.media/catalog/product/apple/iphone/iphone-7/iphone-7-black.jpg",
						},{
							URL: "https://jualku.com/wp-content/uploads/2019/05/apple-iphone-7-2.jpg",
						},
					},
				},{
					Name: "IPhone 7",
					Description: "IPhone 7 Description",
					Price: 700000,
					Discount: 7,
					SoldCount: 100,
					MetaData: "Weight: 0.8gram",
					AddedTime: time.Now(),
					Stock: 70000,
					Rating: 5,
					SubCategoryID:5,
					Images: []*model.ProductImage{
						{
							URL: "https://support.apple.com/library/APPLE/APPLECARE_ALLGEOS/SP705/SP705-iphone_6-mul.png",
						},{
							URL: "https://m.media-amazon.com/images/I/51+K8pddvbS._AC_SX679_.jpg",
						},
					},
				},{
					Name: "IPhone 6",
					Description: "IPhone 6 Description",
					Price: 600000,
					Discount: 6,
					SoldCount: 250,
					MetaData: "Weight: 0.75gram",
					AddedTime: time.Now(),
					Stock: 60000,
					Rating: 5,
					SubCategoryID:5,
					Images: []*model.ProductImage{
						{
							URL: "https://d2xjmi1k71iy2m.cloudfront.net/dairyfarm/id/images/185/0818587_PE774509_S4.jpg",
						},{
							URL: "https://d2xjmi1k71iy2m.cloudfront.net/dairyfarm/id/images/386/0938667_PE794238_S5.jpg",
						},
					},
				},{
					Name: "MacBook Air",
					Description: "MacBook Air Description",
					Price: 8920000,
					Discount: 3,
					SoldCount: 200,
					MetaData: "Weight: 2gram",
					AddedTime: time.Now(),
					Stock: 60000,
					Rating: 5,
					SubCategoryID:6,
					Images: []*model.ProductImage{
						{
							URL: "https://static.bmdstatic.com/pk/product/large/6007d8624ea91.jpg",
						},{
							URL: "https://store.stormfront.co.uk/content/images/thumbs/0014007_macbook_air_blush_pdp_image_position-2_m1_chip_usenjpg.jpeg",
						},
					},
				},
			},
		},{
			UserID: 2,
			Name: "IKEA",
			NameSlug: "IKEA Home Furnishing",
			Points: 70,
			Image: "https://static.au-catalogue-24.com/image/shop/ikea/logo_512.png",
			OperationalStatus: "Open",
			OperationalHour: "07.00 - 21.00",
			Description: "Desc IKEA",
			Slogan: "Slogan IKEA",
			Address: "Alamat IKEA",
			PhoneNumber: "038838832992",
			BadgesID: 2,
			Product: []*model.Product{
				{
					Name: "Sofa",
					Description: "Sofa Description",
					Price: 9200000,
					SoldCount: 300,
					Discount: 10,
					MetaData: "Weight: 10gram",
					AddedTime: time.Now(),
					Stock: 30000,
					Rating: 3,
					SubCategoryID:1,
					Images: []*model.ProductImage{
						{
							URL: "https://d2xjmi1k71iy2m.cloudfront.net/dairyfarm/id/images/185/0818587_PE774509_S4.jpg",
						},{
							URL: "https://d2xjmi1k71iy2m.cloudfront.net/dairyfarm/id/images/386/0938667_PE794238_S5.jpg",
						},
					},
				},{
					Name: "Dining Table",
					Description: "Dining Table Description",
					Price: 7630000,
					Discount: 5,
					SoldCount: 200,
					MetaData: "Weight: 20gram",
					AddedTime: time.Now(),
					Stock: 250000,
					Rating: 4,
					SubCategoryID:2,
					Images: []*model.ProductImage{
						{
							URL: "https://www.ikea.com/us/en/images/products/vedbo-dining-table-black__0815091_pe772752_s5.jpg?f=s",
						},{
							URL: "https://www.ikea.com/us/en/images/products/vedbo-dining-table-black__0766049_pe753697_s5.jpg?f=s",
						},
					},
				}, {
					Name: "Dining Chair",
					Description: "Dining Chair Description",
					Price: 93000,
					Discount: 0,
					MetaData: "Weight: 10gram",
					AddedTime: time.Now(),
					Stock: 5000,
					Rating: 3,
					SoldCount: 250,
					SubCategoryID:1,
					Images: []*model.ProductImage{
						{
							URL: "https://cdn.shopify.com/s/files/1/2350/5189/products/Costa_dining_chair_side_new_800x.jpg?v=1580267464",
						},{
							URL: "https://cdn.shopify.com/s/files/1/2350/5189/products/Costa_dining_chair_new.jpg?v=1580267381",
						},
					},
				}, {
					Name: "Cafe Table",
					Description: "Cafe Table Description",
					Price: 330000,
					Discount: 5,
					MetaData: "Weight: 20gram",
					AddedTime: time.Now(),
					Stock: 250000,
					Rating: 4,
					SoldCount: 250,
					SubCategoryID:2,
					Images: []*model.ProductImage{
						{
							URL: "https://media.fds.fi/product_image/800/409Muuto_AK.jpg",
						},{
							URL: "https://homeexporter.com/wp-content/uploads/2021/06/cafe-table-500x500-1.jpg",
						},
					},
				},				{
					Name: "Kids Chair",
					Description: "Kids Chair Description",
					Price: 770000,
					Discount: 15,
					MetaData: "Weight: 5gram",
					AddedTime: time.Now(),
					Stock: 4000,
					Rating: 5,
					SoldCount: 250,
					SubCategoryID:1,
					Images: []*model.ProductImage{
						{
							URL: "https://i5.walmartimages.com/asr/f25bcbd9-edb2-497a-8349-33e4149fe92f_1.11b6a37662cf1e9353b51b11eb49e478.jpeg?odnHeight=612&odnWidth=612&odnBg=FFFFFF",
						},{
							URL: "https://m.media-amazon.com/images/I/71bIrPF9ZOL._SX569_.jpg",
						},
					},
				},{
					Name: "Kids Table",
					Description: "Kids Table Description",
					Price: 230000,
					Discount: 13,
					MetaData: "Weight: 5gram",
					AddedTime: time.Now(),
					Stock: 250000,
					Rating: 5,
					SoldCount: 250,
					SubCategoryID:2,
					Images: []*model.ProductImage{
						{
							URL: "https://www.ikea.com/us/en/images/products/mammut-childrens-table-indoor-outdoor-blue__0735844_pe740211_s5.jpg?f=xs",
						},{
							URL: "https://www.ikea.com/us/en/images/products/mammut-childrens-table-indoor-outdoor-red__0735839_pe740209_s5.jpg?f=xs",
						},
					},
				},
			},
		},{
			UserID: 3,
			Name: "BeautyLab",
			NameSlug: "BeautyLab Official",
			Points: 110,
			Image: "https://1.bp.blogspot.com/-qb6SULBLCZo/X36Ujrsw9xI/AAAAAAAAAzk/-a5NX3xr2t41q3iGz64DdEpCR7M-wul-gCLcBGAsYHQ/s2048/1602121956049.png",
			OperationalStatus: "Open",
			OperationalHour: "06.00 - 15.00",
			Description: "Desc BeautyLab",
			Slogan: "Slogan BeautyLab",
			Address: "Alamat BeautyLab",
			PhoneNumber: "0283839923",
			BadgesID: 3,
			Product: []*model.Product{
				{
					Name: "Body Scrub",
					Description: "Body Scrub Description",
					Price: 55000,
					Discount: 5,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 10000,
					SoldCount: 250,
					Rating: 5,
					SubCategoryID:13,
					Images: []*model.ProductImage{
						{
							URL: "https://cf.shopee.co.id/file/b912a4e15491159e041507b0e2b7bfec",
						},{
							URL: "https://cf.shopee.co.id/file/e0752cfa1502d57d63ff78e7c2b0fb0a",
						},
					},
				},{
					Name: "Shower Scrub",
					Description: "Shower Scrub Description",
					Price: 660000,
					Discount: 20,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 200000,
					SoldCount: 300,
					Rating: 4,
					SubCategoryID:13,
					Images: []*model.ProductImage{
						{
							URL: "https://cf.shopee.co.id/file/279200f78807db8d820adee6f77fc54c",
						},{
							URL: "https://image.femaledaily.com/dyn/640/images/prod-pics/product_1603968990_Scarlett_800x800.jpg",
						},
					},
				},{
					Name: "Body Lotion",
					Description: "Body Lotion Description",
					Price: 600000,
					Discount: 30,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 200000,
					SoldCount: 700,
					Rating: 5,
					SubCategoryID:13,
					Images: []*model.ProductImage{
						{
							URL: "https://id-live-05.slatic.net/shop/0b6ac54b74ae246ecdd4a28cbb4737fd.jpeg_2200x2200q80.jpg_.webp",
						},{
							URL: "https://cf.shopee.co.id/file/3adee1925031b6735c47e38aa10a1dba",
						},
					},
				},{
					Name: "Hair Serum",
					Description: "Hair Serum Description",
					Price: 6600,
					Discount: 40,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 50000,
					Rating: 4,
					SoldCount: 1000,
					SubCategoryID:14,
					Images: []*model.ProductImage{
						{
							URL: "https://cdn.idntimes.com/content-images/post/20211112/51ptxufjil-sl1080-2ad592db4fb55187edca9d98c4f61fe1.jpg",
						},{
							URL: "https://cf.shopee.co.id/file/17505480c80c7745ea106a2da2fd84d7",
						},
					},
				},{
					Name: "Kuteks",
					Description: "Kuteks Description",
					Price: 440000,
					Discount: 70,
					MetaData: "Weight: 1gram",
					AddedTime: time.Now(),
					Stock: 70000,
					Rating: 5,
					SoldCount: 10,
					SubCategoryID:15,
					Images: []*model.ProductImage{
						{
							URL: "https://cf.shopee.co.id/file/b84a884992b3bba4afa6f9883ae0718a",
						},{
							URL: "https://images.tokopedia.net/img/cache/700/VqbcmM/2020/12/4/7861e2c4-7408-43bc-a069-aece7827a878.jpg",
						},
					},
				},
			},
		},
	}

	shopVendor := []model.ShopShippingVendor{
		{
			ShopID: 1,
			VendorID: 1,
		},{
			ShopID: 1,
			VendorID: 2,
		},{
			ShopID: 1,
			VendorID: 3,
		},{
			ShopID: 2,
			VendorID: 3,
		},{
			ShopID: 2,
			VendorID: 4,
		},
	}

	shopVoucher := []model.ShopVoucher{
		{
			ShopID: 1,
			VoucherID: 1,
		},{
			ShopID: 1,
			VoucherID: 2,
		},{
			ShopID: 1,
			VoucherID: 3,
		},{
			ShopID: 2,
			VoucherID: 3,
		},{
			ShopID: 2,
			VoucherID: 1,
		},
	}

	campaign := []model.Campaign{
		{
			URL: "https://ecs7.tokopedia.net/img/cache/730/kjjBfF/2021/12/2/5b5ec00c-09c9-48cb-b295-c9d3eeae6cf9.jpg",
		}, {
			URL: "https://ecs7.tokopedia.net/img/kjjBfF/2021/1/28/79fec4a4-8220-45ca-8593-8b18eba14ffc.png",
		}, {
			URL: "https://ecs7.tokopedia.net/img/cache/730/kjjBfF/2021/5/19/f446d9b2-557d-4ae4-bbde-0e8eee846635.png",
		}, {
			URL: "https://ecs7.tokopedia.net/blog-tokopedia-com/uploads/2020/07/Tokopedia-2-3.jpg",
		}, {
			URL: "https://ecs7.tokopedia.net/blog-tokopedia-com/uploads/2021/05/Banner_Gosend-Tokopedia-Campaign.jpg",
		}, {
			URL: "https://lelogama.go-jek.com/post_thumbnail/Copy_of_Tokped_1456x818.jpg",
		}, {
			URL: "https://cdn.promolist.id/promo/1/611b5e84d29c6_d4f268dd28d20e61ec73f7e030059c40714e1ae60b2773e199ad35db93bcc14c.jpg",
		}, {
			URL: "https://1.bp.blogspot.com/-rl2PpwU1ZXw/YV5qMaBmNrI/AAAAAAAAMuc/QeF-aq4x2oM2PKvtrdAEyk2xvx2a77Z_wCNcBGAsYHQ/s800/pulsa%2BTokopedia.jpg",
		}, {
			URL: "https://ecs7.tokopedia.net/blog-tokopedia-com/uploads/2021/01/Tokopedia-x-BLACKPINK-1-1024x512.jpg",
		}, {
			URL: "https://siaranpers.id/wp-content/uploads/2021/01/Tokopedia-x-BTS-1.jpg",
		}, {
			URL: "https://ecs7.tokopedia.net/blog-tokopedia-com/uploads/2019/10/Tokopedia-Blog_BTS-Campaign_1200x620-1024x529.jpg",
		}, {
			URL: "https://pbs.twimg.com/media/EF7pxQ0UcAETWZC.jpg",
		}, {
			URL: "https://images.tokopedia.net/img/Template/FB-Twitter-Mega-cashback-Oktober.jpg",
		}, {
			URL: "http://jagoanindonesia.id/wp-content/uploads/2020/06/tokopedia-toppicks_1024x1024-1.jpg",
		}, {
			URL: "https://ecs7.tokopedia.net/img/blog/seller/2020/09/statistik-new-768x400.jpg",
		}, 
	}
	db.Create(&campaign)
	db.Create(&badge)
	db.Create(&shipping)
	db.Create(&voucher)
	db.Create(&category)
	db.Create(&user)
	db.Create(&subCategory)
	db.Create(&shop)
	db.Create(&shopVendor)
	db.Create(&shopVoucher)
}

