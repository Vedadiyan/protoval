# Protoval 
Protoval is a dynamic protobuf validation library for Golang that validates protobuf fields, lists, and objects at runtime. Protoval is simple to use and does not require any dependency on any `.proto` file. It can be used out of the box without any configurations. 

## Usage 
All you should do is to define a field option on your `.proto` file with any given name. For example: 

    syntax = "proto3";

    import "google/protobuf/descriptor.proto";

    extend google.protobuf.FieldOptions {
        string validate = 10000;
    }

Then you can use this field option to validate your fields. For example: 

    message Test {
        string fieldA = 1 [(validate) = "required|min_len:10|max_len:100"];
    }

In your Golang code, you can invoke your validation pattern by calling the `Validate` method. For example: 

    validationContext := protoval.New("validate", SomeMessage)
    err := validationContext.Validate()
    if err != nil {
        // Do something about the error
    }
    if !validationContext.IsValid() {
        for _, err := range validationContext.Errors() {
            fmt.Println(err)
        }
    }

You can register any validation function by calling the `Register` method. For example: 

    protoval.Register("myValidationMethod", func (name string, value any, rule string) error {
        if value == nil {
            return protoval.NewError(name, "is required")
        }
        return nil
    })

You can also use Regex expressions out of the box by encapsulating them inside `/` character. For example: 

    message Test {
        string fieldA = 1 [(validate) = "required|min_len:10|max_len:100|/[Aa-Zz]*/"];
    }

Regex expression can appear anywhere in the pattern and NOT only at the end. 